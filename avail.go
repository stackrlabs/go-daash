package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"log"

	"encoding/binary"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/rollkit/go-da"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

type DAClient interface {
	PostData(txData []byte) (*BatchDAData, error)
}

type AccountNextIndexRPCResponse struct {
	Result uint `json:"result"`
}

type DataProofRPCResponse struct {
	Result DataProof `json:"result"`
}
type DataProof struct {
	Root           string   `json:"root"`
	Proof          []string `json:"proof"`
	NumberOfLeaves uint32   `json:"numberOfLeaves"`
	LeafIndex      uint32   `json:"leafIndex"`
	Leaf           string   `json:"leaf"`
}

type AvailDA struct {
	config             Config
	API                *gsrpc.SubstrateAPI
	Meta               *types.Metadata
	AppID              int
	GenesisHash        types.Hash
	Rv                 *types.RuntimeVersion
	KeyringPair        signature.KeyringPair
	DestinationAddress types.Hash
	DestinationDomain  types.UCompact
}

func NewAvailDA() (*AvailDA, error) {
	a := AvailDA{}
	err := a.config.GetConfig("./avail-config.json")
	if err != nil {
		return nil, fmt.Errorf("cannot get config", err)
	}

	a.API, err = gsrpc.NewSubstrateAPI(a.config.APIURL)
	if err != nil {
		// log.Error("cannot get api:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get api", err)
	}

	a.Meta, err = a.API.RPC.State.GetMetadataLatest()
	if err != nil {
		// log.Error("cannot get metadata:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get metadata", err)
	}

	a.AppID = 0

	// if app id is greater than 0 then it must be created before submitting data
	if a.config.AppID != 0 {
		a.AppID = a.config.AppID
	}

	a.GenesisHash, err = a.API.RPC.Chain.GetBlockHash(0)
	if err != nil {
		// log.Error("cannot get genesis hash:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get genesis hash", err)
	}

	a.Rv, err = a.API.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		// log.Error("cannot get runtime version:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get runtime version", err)
	}

	a.KeyringPair, err = signature.KeyringPairFromSecret(a.config.Seed, 42)
	if err != nil {
		// log.Error("cannot get keyring pair:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get keyring pair", err)
	}

	// DestinationAddress, err = types.NewHashFromHexString(Config.DestinationAddress)
	// if err != nil {
	// 	log.Fatalf("cannot decode destination address:%w", err)
	// }

	// DestinationDomain = types.NewUCompactFromUInt(uint64(Config.DestinationDomain))
	fmt.Println("🟢 Avail DA initialized", a)
	return &a, nil
}

// MaxBlobSize returns the max blob size
func (c *AvailDA) MaxBlobSize(ctx context.Context) (uint64, error) {
	var maxBlobSize uint64 = 64 * 64 * 500
	return maxBlobSize, nil
}

// Submit a list of blobs to Avail DA
// Currently, we submit to a trusted RPC Avail node. In the future, we will submit via an Avail light client.
func (a *AvailDA) Submit(ctx context.Context, daBlobs []da.Blob, gasPrice float64) ([]da.ID, []da.Proof, error) {
	// TODO: Add support for multiple blobs
	log.Println("data", zap.Any("data", daBlobs[0]))
	log.Println("⚡️ Preparing to post data to Avail:%d bytes", zap.Int("data_size", len(daBlobs[0])))
	fmt.Println(*a)
	newCall, err := types.NewCall(a.Meta, "DataAvailability.submit_data", types.NewBytes(daBlobs[0]))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create new call", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := a.GetAccountNextIndex()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get account next index", err)
	}

	options := types.SignatureOptions{
		BlockHash:          a.GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        a.GenesisHash,
		Nonce:              nonce,
		SpecVersion:        a.Rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(1000),
		AppID:              types.NewUCompactFromUInt(uint64(a.AppID)),
		TransactionVersion: a.Rv.TransactionVersion,
	}

	err = ext.Sign(a.KeyringPair, options)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot sign extrinsic", err)
	}

	// Send the extrinsic
	sub, err := a.API.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot submit extrinsic", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(a.config.Timeout) * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Println("📥 Submit data extrinsic included in block %v", zap.String("status in block", status.AsInBlock.Hex()))
			}
			if status.IsFinalized {
				blockHash = status.AsFinalized
				break out
			} else if status.IsDropped {
				return nil, nil, fmt.Errorf("extrinsic dropped")
			} else if status.IsUsurped {
				return nil, nil, fmt.Errorf("extrinsic usurped")
			} else if status.IsRetracted {
				return nil, nil, fmt.Errorf("extrinsic retracted")
			} else if status.IsInvalid {
				return nil, nil, fmt.Errorf("extrinsic invalid")
			}
		case <-timeout:
			return nil, nil, fmt.Errorf("timeout")
		}
	}

	log.Println("✅ Data submitted by sequencer bytes against AppID sent with hash", zap.String("block hash", blockHash.Hex()))

	var dataProof DataProof
	var batchHash [32]byte

	h := sha3.NewLegacyKeccak256()
	h.Write(daBlobs[0])
	h.Sum(batchHash[:0])

	block, err := a.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get block", err)
	}

	// TODO: We go through all extrinsics to get the data proof. We should be able to get the data proof from the extrinsic that we submitted.
	for i := 1; i <= len(block.Block.Extrinsics); i++ {
		resp, err := http.Post("https://goldberg.avail.tools/api", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", i, blockHash))) //nolint: noctx
		if err != nil {
			return nil, nil, fmt.Errorf("cannot post query request", err)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot read body", err)
		}
		err = resp.Body.Close()
		if err != nil {
			return nil, nil, fmt.Errorf("cannot close body", err)
		}

		var dataProofResp DataProofRPCResponse
		err = json.Unmarshal(data, &dataProofResp)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot unmarshal data proof", err)
		}

		if dataProofResp.Result.Leaf == fmt.Sprintf("%#x", batchHash) {
			dataProof = dataProofResp.Result
			break
		}
	}
	// NOTE: Substrate's BlockNumber type is an alias for u32 type, which is uint32
	blobID := a.makeID(uint32(block.Block.Header.Number), dataProof.LeafIndex)
	blobIDs := make([]da.ID, 1)
	blobIDs[0] = blobID

	serialisedProofs := make([]da.Proof, 1)
	for _, word := range dataProof.Proof {
		serialisedProofs[0] = append(serialisedProofs[0], word...)
	}

	log.Println("💿 received data proof:%+v", zap.Any("dataproof", dataProof))
	return blobIDs, serialisedProofs, nil
}

// Get returns Blob for each given ID, or an error.
func (a *AvailDA) Get(ctx context.Context, ids []da.ID) ([]da.Blob, error) {
	// TODO: We are dealing with single blobs for now. We will need to handle multiple blobs in the future.
	blockHeight, leafIndex := a.splitID(ids[0])
	data, err := a.GetData(uint64(blockHeight), uint(leafIndex))
	if err != nil {
		return nil, fmt.Errorf("cannot get data", err)
	}
	log.Println("📥 received data:%+v", zap.Any("data", data))
	return []da.Blob{data}, nil
}

// GetIDs returns IDs of all Blobs located in DA at given height.
func (a *AvailDA) GetIDs(ctx context.Context, height uint64) ([]da.ID, error) {
	// TODO: Need to implement this
	return nil, nil
}

// Commit creates a Commitment for each given Blob.
func (a *AvailDA) Commit(ctx context.Context, daBlobs []da.Blob) ([]da.Commitment, error) {
	// TODO: Need to implement this
	return nil, nil
}

// Validate validates Commitments against the corresponding Proofs. This should be possible without retrieving the Blobs.
func (c *AvailDA) Validate(ctx context.Context, ids []da.ID, daProofs []da.Proof) ([]bool, error) {
	// TODO: Need to implement this
	return nil, nil
}

func (a AvailDA) PostData(txData []byte) (*BatchDAData, error) {
	log.Println("⚡️ Prepared data for Avail:%d bytes", zap.Int("data_size", len(txData)))

	newCall, err := types.NewCall(a.Meta, "DataAvailability.submit_data", types.NewBytes(txData))
	if err != nil {
		return nil, fmt.Errorf("cannot create new call", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := a.GetAccountNextIndex()
	if err != nil {
		return nil, fmt.Errorf("cannot get account next index", err)
	}

	options := types.SignatureOptions{
		BlockHash:          a.GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        a.GenesisHash,
		Nonce:              nonce,
		SpecVersion:        a.Rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(1000),
		AppID:              types.NewUCompactFromUInt(uint64(a.AppID)),
		TransactionVersion: a.Rv.TransactionVersion,
	}

	err = ext.Sign(a.KeyringPair, options)
	if err != nil {
		return nil, fmt.Errorf("cannot sign extrinsic", err)
	}

	// Send the extrinsic
	sub, err := a.API.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, fmt.Errorf("cannot submit extrinsic", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(a.config.Timeout) * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Println("📥 Submit data extrinsic included in block %v", zap.String("status in block", status.AsInBlock.Hex()))
			}
			if status.IsFinalized {
				blockHash = status.AsFinalized
				break out
			} else if status.IsDropped {
				return nil, fmt.Errorf("extrinsic dropped")
			} else if status.IsUsurped {
				return nil, fmt.Errorf("extrinsic usurped")
			} else if status.IsRetracted {
				return nil, fmt.Errorf("extrinsic retracted")
			} else if status.IsInvalid {
				return nil, fmt.Errorf("extrinsic invalid")
			}
		case <-timeout:
			return nil, fmt.Errorf("timeout")
		}
	}

	log.Println("✅ Data submitted by sequencer bytes against AppID sent with hash", zap.String("block hash", blockHash.Hex()))

	var dataProof DataProof
	var batchHash [32]byte

	h := sha3.NewLegacyKeccak256()
	h.Write(txData)
	h.Sum(batchHash[:0])

	block, err := a.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block", err)
	}

	for i := 1; i <= len(block.Block.Extrinsics); i++ {
		resp, err := http.Post("https://goldberg.avail.tools/api", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", i, blockHash))) //nolint: noctx
		if err != nil {
			return nil, fmt.Errorf("cannot post query request", err)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("cannot read body", err)
		}
		err = resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("cannot close body", err)
		}

		fmt.Println("data", data)
		var dataProofResp DataProofRPCResponse
		err = json.Unmarshal(data, &dataProofResp)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal data proof", err)
		}

		if dataProofResp.Result.Leaf == fmt.Sprintf("%#x", batchHash) {
			dataProof = dataProofResp.Result
			break
		}
	}

	log.Println("💿 received data proof:%+v", zap.Any("dataproof", dataProof))
	var batchDAData BatchDAData
	batchDAData.Proof = dataProof.Proof
	batchDAData.Width = uint(dataProof.NumberOfLeaves)
	batchDAData.LeafIndex = uint(dataProof.LeafIndex)

	header, err := a.API.RPC.Chain.GetHeader(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get header", err)
	}

	batchDAData.BlockNumber = uint(header.Number)
	data, err := a.GetData(uint64(header.Number), uint(dataProof.LeafIndex))
	if err != nil {
		return nil, fmt.Errorf("cannot get data", err)
	}
	log.Println("📥 received data:%+v", zap.Any("data", data))
	log.Println("🟢 prepared DA data:%+v", zap.Any("batchdata", batchDAData))

	return &batchDAData, nil
}

type BatchDAData struct {
	BlockNumber uint
	Proof       []string `json:"proof"`
	Width       uint     `json:"number_of_leaves"`
	LeafIndex   uint     `json:"leaf_index"`
}

func (b BatchDAData) IsEmpty() bool {
	return reflect.DeepEqual(b, BatchDAData{})
}

func (a AvailDA) GetAccountNextIndex() (types.UCompact, error) {
	// TODO: Add context to the request
	resp, err := http.Post("https://goldberg.avail.tools/api", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"system_accountNextIndex\",\"params\":[\"%v\"]}", a.KeyringPair.Address))) //nolint: noctx
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot post account next index request", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot read body", err)
	}

	var accountNextIndex AccountNextIndexRPCResponse
	err = json.Unmarshal(data, &accountNextIndex)
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot unmarshal account next index", err)
	}

	return types.NewUCompactFromUInt(uint64(accountNextIndex.Result)), nil
}

// makeID creates a unique ID to reference a blob on Avail
func (a *AvailDA) makeID(blockHeight uint32, leafIndex uint32) da.ID {
	// Serialise height and leaf index to binary
	heightLen := 4
	leafIndexLen := 4
	heightBytes := make([]byte, heightLen)
	leafIndexBytes := make([]byte, leafIndexLen)
	binary.LittleEndian.PutUint32(heightBytes, blockHeight)
	binary.LittleEndian.PutUint32(leafIndexBytes, leafIndex)
	idBytes := append(heightBytes, leafIndexBytes...)
	return da.ID(idBytes)
}

// splitID returns the block height and leaf index from a unique ID
func (a *AvailDA) splitID(id da.ID) (uint32, uint32) {
	heightLen := 4
	leafIndexLen := 4
	heightBytes := id[:heightLen]
	leafIndexBytes := id[heightLen : heightLen+leafIndexLen]
	blockHeight := binary.LittleEndian.Uint32(heightBytes)
	leafIndex := binary.LittleEndian.Uint32(leafIndexBytes)
	return blockHeight, leafIndex
}

func (a AvailDA) GetData(blockNumber uint64, index uint) ([]byte, error) {
	blockHash, err := a.API.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		return nil, fmt.Errorf("cannot get block hash", err)
	}

	block, err := a.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block", err)
	}

	var data [][]byte
	for _, ext := range block.Block.Extrinsics {
		if ext.Method.CallIndex.SectionIndex == 29 && ext.Method.CallIndex.MethodIndex == 1 {
			data = append(data, ext.Method.Args[2:])
		}
	}

	return data[index], nil
}

type Config struct {
	Seed               string `json:"seed"`
	APIURL             string `json:"api_url"`
	AppID              int    `json:"app_id"`
	DestinationDomain  int    `json:"destination_domain"`
	DestinationAddress string `json:"destination_address"`
	Timeout            int    `json:"timeout"`
}

func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return fmt.Errorf("cannot open config file", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("cannot read config file", err)
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file", err)
	}

	return nil
}
