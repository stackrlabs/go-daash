package avail

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

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stackrlabs/go-daash/da"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
)

type AccountNextIndexRPCResponse struct {
	Result uint `json:"result"`
}
type DataProofRPCResponse struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		DataProof `json:"dataProof"`
	} `json:"result"`
}
type DataProof struct {
	Leaf           string   `json:"leaf"`
	LeafIndex      int64    `json:"leafIndex"`
	NumberOfLeaves int64    `json:"numberOfLeaves"`
	Proof          []string `json:"proof"`
	Roots          struct {
		BlobRoot   string `json:"blobRoot"`
		BridgeRoot string `json:"bridgeRoot"`
		DataRoot   string `json:"dataRoot"`
	} `json:"roots"`
}

type Client struct {
	Config             Config
	API                *gsrpc.SubstrateAPI
	Meta               *types.Metadata
	AppID              int
	GenesisHash        types.Hash
	Rv                 *types.RuntimeVersion
	KeyringPair        signature.KeyringPair
	DestinationAddress types.Hash
	DestinationDomain  types.UCompact
}

// Returns a newly initalised Avail DA client
func NewClient(configPath string) (*Client, error) {
	a := Client{}
	err := a.Config.GetConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot get config:%w", err)
	}

	a.API, err = gsrpc.NewSubstrateAPI(a.Config.WsRpcURL)
	if err != nil {
		// log.Error("cannot get api:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get api:%w", err)
	}

	a.Meta, err = a.API.RPC.State.GetMetadataLatest()
	if err != nil {
		// log.Error("cannot get metadata:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get metadata:%w", err)
	}

	a.AppID = 0

	// if app id is greater than 0 then it must be created before submitting data
	if a.Config.AppID != 0 {
		a.AppID = a.Config.AppID
	}

	a.GenesisHash, err = a.API.RPC.Chain.GetBlockHash(0)
	if err != nil {
		// log.Error("cannot get genesis hash:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get genesis hash:%w", err)
	}

	a.Rv, err = a.API.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		// log.Error("cannot get runtime version:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get runtime version:%w", err)
	}

	a.KeyringPair, err = signature.KeyringPairFromSecret(a.Config.Seed, 42)
	if err != nil {
		// log.Error("cannot get keyring pair:%w", zap.Error(err))
		return nil, fmt.Errorf("cannot get keyring pair:%w", err)
	}

	return &a, nil
}

// MaxBlobSize returns the max blob size
func (c *Client) MaxBlobSize(ctx context.Context) (uint64, error) {
	var maxBlobSize uint64 = 64 * 64 * 500
	return maxBlobSize, nil
}

// Submit a list of blobs to Avail DA
// Currently, we submit to a trusted RPC Avail node. In the future, we will submit via an Avail light client.
func (a *Client) Submit(ctx context.Context, daBlobs []da.Blob, gasPrice float64) ([]da.ID, []da.Proof, error) {
	// TODO: Add support for multiple blobs
	daBlob := daBlobs[0]
	log.Println("data", zap.String("data", string(daBlob)))
	log.Printf("⚡️ Preparing to post data to Avail:%d bytes", len(daBlob))
	newCall, err := types.NewCall(a.Meta, "DataAvailability.submit_data", types.NewBytes(daBlob))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := a.GetAccountNextIndex()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get account next index:%w", err)
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

	fmt.Println("options transaction version", options.TransactionVersion, "spec version", options.SpecVersion)

	err = ext.Sign(a.KeyringPair, options)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot sign extrinsic:%w", err)
	}

	// Send the extrinsic
	sub, err := a.API.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot submit extrinsic:%w", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(a.Config.Timeout) * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Printf("📥 Submit data extrinsic included in block %v", status.AsInBlock.Hex())
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

	var batchHash [32]byte

	h := sha3.NewLegacyKeccak256()
	h.Write(daBlobs[0])
	h.Sum(batchHash[:0])

	block, err := a.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get block:%w", err)
	}

	var dataProofResp DataProofRPCResponse
	var extIndex int
	for idx, e := range block.Block.Extrinsics {
		// Look for our submitted extrinsic in the block
		extBytes, err := json.Marshal(ext)
		if err != nil {
			continue
		}
		extBytes = []byte(strings.Trim(string(extBytes), "\""))

		eBytes, err := json.Marshal(e)
		if err != nil {
			continue
		}
		eBytes = []byte(strings.Trim(string(eBytes), "\""))
		if string(extBytes) == string(eBytes) {
			extIndex = idx
			resp, err := http.Post(a.Config.HttpApiURL, "application/json",
				strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProofV2\",\"params\":[%d, \"%#x\"]}", idx+1, blockHash))) //nolint: noctx
			if err != nil {
				break
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				break
			}
			err = resp.Body.Close()
			if err != nil {
				break
			}
			err = json.Unmarshal(data, &dataProofResp)
			if err != nil {
				break
			}

			if dataProofResp.Result.DataProof.Leaf == fmt.Sprintf("%#x", batchHash) {
				log.Println("data proof validated", dataProofResp.Result)
			}
			break
		}
	}
	dataProof := dataProofResp.Result.DataProof
	// NOTE: Substrate's BlockNumber type is an alias for u32 type, which is uint32
	blobID := ID{Height: uint64(block.Block.Header.Number), ExtIndex: uint32(extIndex)}
	blobIDs := make([]da.ID, 1)
	blobIDs[0] = blobID

	serialisedProofs := make([]da.Proof, 1)
	for _, word := range dataProof.Proof {
		serialisedProofs[0] = append(serialisedProofs[0], word...)
	}

	log.Printf("💿 received data proof:%+v\n", zap.Any("dataproof", dataProof))
	return blobIDs, serialisedProofs, nil
}

// Get returns Blob for each given ID, or an error.
func (a *Client) Get(ctx context.Context, ids []da.ID) ([]da.Blob, error) {
	// TODO: We are dealing with single blobs for now. We will need to handle multiple blobs in the future.
	ext, err := a.GetExtrinsic(ids[0])
	if err != nil {
		return nil, fmt.Errorf("cannot get extrinsic:%w", err)
	}
	blobData := ext.Method.Args[2:]
	log.Printf("📥 received data:%+v", blobData)
	return []da.Blob{blobData}, nil
}

// GetIDs returns IDs of all Blobs located in DA at given height.
func (a *Client) GetIDs(ctx context.Context, height uint64) ([]da.ID, error) {
	// TODO: Need to implement this
	return nil, nil
}

// Commit creates a Commitment for each given Blob.
func (a *Client) Commit(ctx context.Context, daBlobs []da.Blob) ([]da.Commitment, error) {
	// TODO: Need to implement this
	return nil, nil
}

// GetProofs returns the proofs for the given IDs
func (a *Client) GetProof(ctx context.Context, blockHeight uint32, extIdx int) (DataProofRPCResponse, error) {
	var dataProofResp DataProofRPCResponse
	blockHash, err := a.API.RPC.Chain.GetBlockHash(uint64(blockHeight))
	if err != nil {
		return dataProofResp, fmt.Errorf("cannot get block hash:%w", err)
	}
	resp, err := http.Post(a.Config.HttpApiURL, "application/json",
		strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", extIdx, blockHash)))

	if err != nil {
		return dataProofResp, fmt.Errorf("cannot get data proof:%w", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return dataProofResp, fmt.Errorf("cannot read data:%w", err)
	}
	err = resp.Body.Close()
	if err != nil {
		return dataProofResp, fmt.Errorf("cannot close body:%w", err)
	}
	fmt.Println("raw proof data", string(data))
	err = json.Unmarshal(data, &dataProofResp)
	if err != nil {
		return dataProofResp, fmt.Errorf("cannot unmarshal data:%w", err)
	}
	fmt.Println("dataProofResp", dataProofResp)
	return dataProofResp, nil
}

// Validate validates Commitments against the corresponding Proofs. This should be possible without retrieving the Blobs.
func (c *Client) Validate(ctx context.Context, ids []da.ID, daProofs []da.Proof) ([]bool, error) {
	// TODO: Need to implement this
	return nil, nil
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

func (a *Client) GetAccountNextIndex() (types.UCompact, error) {
	// TODO: Add context to the request
	resp, err := http.Post(a.Config.HttpApiURL, "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"system_accountNextIndex\",\"params\":[\"%v\"]}", a.KeyringPair.Address))) //nolint: noctx
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot post account next index request:%w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot read body:%w", err)
	}
	var accountNextIndex AccountNextIndexRPCResponse
	err = json.Unmarshal(data, &accountNextIndex)
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot unmarshal account next index:%w Response: %s", err, data)
	}

	return types.NewUCompactFromUInt(uint64(accountNextIndex.Result)), nil
}

type ID struct {
	Height   uint64 `json:"blockHeight"`
	ExtIndex uint32 `json:"extIdx"`
}

type Config struct {
	Seed               string `json:"seed"`
	WsRpcURL           string `json:"wsRpcUrl"`
	HttpApiURL         string `json:"httpApiUrl"`
	AppID              int    `json:"app_id"`
	DestinationDomain  int    `json:"destination_domain"`
	DestinationAddress string `json:"destination_address"`
	Timeout            int    `json:"timeout"`
	Network            string `json:"network"`
}

func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return fmt.Errorf("cannot open config file:%w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("cannot read config file:%w", err)
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config file:%w", err)
	}

	return nil
}

func (a *Client) GetExtrinsic(id da.ID) (types.Extrinsic, error) {
	availID, ok := id.(ID)
	if !ok {
		return types.Extrinsic{}, fmt.Errorf("invalid ID")
	}
	blockHash, err := a.API.RPC.Chain.GetBlockHash(uint64(availID.Height))
	if err != nil {
		return types.Extrinsic{}, fmt.Errorf("cannot get block hash:%w", err)
	}
	block, err := a.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return types.Extrinsic{}, fmt.Errorf("cannot get block:%w", err)
	}
	return block.Block.Extrinsics[availID.ExtIndex], nil
}
