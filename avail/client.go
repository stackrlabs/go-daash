package avail

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
func (c *Client) Submit(ctx context.Context, daBlob da.Blob, gasPrice float64) (da.ID, error) {
	log.Printf("⚡️ Preparing to post data to Avail:%d bytes", len(daBlob))
	newCall, err := types.NewCall(c.Meta, "DataAvailability.submit_data", types.NewBytes(daBlob))
	if err != nil {
		return nil, fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := c.GetAccountNextIndex()
	if err != nil {
		return nil, fmt.Errorf("cannot get account next index:%w", err)
	}

	options := types.SignatureOptions{
		BlockHash:          c.GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.GenesisHash,
		Nonce:              nonce,
		SpecVersion:        c.Rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(1000),
		AppID:              types.NewUCompactFromUInt(uint64(c.AppID)),
		TransactionVersion: c.Rv.TransactionVersion,
	}

	err = ext.Sign(c.KeyringPair, options)
	if err != nil {
		return nil, fmt.Errorf("cannot sign extrinsic:%w", err)
	}

	// Send the extrinsic
	sub, err := c.API.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, fmt.Errorf("cannot submit extrinsic:%w", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(c.Config.Timeout) * time.Second)
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

	block, err := c.API.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block:%w", err)
	}

	var extIndex int
	// Right now, we look trhough all extrinsics in the block to find our submitted extrinsic's index
	// This works for now but we should find a better way to do this (TODO)
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
		}
	}

	blobHash, err := c.Commit(ctx, daBlob)
	if err != nil {
		return nil, fmt.Errorf("cannot commit:%w", err)
	}
	// NOTE: Substrate's BlockNumber type is an alias for u32 type, which is uint32
	blobID := ID{
		Height:   uint64(block.Block.Header.Number),
		ExtIndex: uint32(extIndex),
		BlobHash: blobHash.([32]byte),
	}

	return blobID, nil
}

// Get returns Blob for each given ID, or an error.
func (c *Client) Get(ctx context.Context, id da.ID) (da.Blob, error) {
	// TODO: We are dealing with single blobs for now. We will need to handle multiple blobs in the future.
	ext, err := c.GetExtrinsic(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get extrinsic:%w", err)
	}
	blobData := ext.Method.Args[2:]
	log.Printf("📥 received data:%+v", blobData)
	return blobData, nil
}

// Commit creates a Commitment for a given Blob.
func (a *Client) Commit(ctx context.Context, daBlob da.Blob) (da.Commitment, error) {
	var blobHash [32]byte
	h := sha3.NewLegacyKeccak256()
	h.Write(daBlob)
	h.Sum(blobHash[:0])

	return blobHash, nil
}

// GetProof returns the proof of inclusion for the given ID
func (a *Client) GetProof(ctx context.Context, id da.ID) (da.Proof, error) {
	availID, ok := id.(ID)
	if !ok {
		return nil, fmt.Errorf("invalid ID")
	}
	var dataProofResp DataProofRPCResponse
	blockHash, err := a.API.RPC.Chain.GetBlockHash(uint64(availID.Height))
	if err != nil {
		return dataProofResp, fmt.Errorf("cannot get block hash:%w", err)
	}
	resp, err := http.Post(a.Config.HttpApiURL, "application/json",
		strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", availID.ExtIndex, blockHash)))

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
func (c *Client) Validate(ctx context.Context, id da.ID, daProof da.Proof) (bool, error) {
	availID, ok := id.(ID)
	if !ok {
		return false, fmt.Errorf("invalid ID")
	}
	proof, ok := daProof.(Proof)
	if !ok {
		return false, fmt.Errorf("invalid proof")
	}

	return proof.Result.DataProof.Leaf == fmt.Sprintf("%#x", availID.BlobHash), nil
}

// Utility functions
type AccountNextIndexRPCResponse struct {
	Result uint `json:"result"`
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
