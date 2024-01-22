package main

import (
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
	NumberOfLeaves uint     `json:"numberOfLeaves"`
	LeafIndex      uint     `json:"leafIndex"`
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

func NewAvailDA() *AvailDA {
	a := AvailDA{}
	err := a.config.GetConfig("./avail-config.json")
	if err != nil {
		// log.Error("cannot get config:%w", zap.Error(err))
		return nil
	}

	a.API, err = gsrpc.NewSubstrateAPI(a.config.APIURL)
	if err != nil {
		// log.Error("cannot get api:%w", zap.Error(err))
		return nil
	}

	a.Meta, err = a.API.RPC.State.GetMetadataLatest()
	if err != nil {
		// log.Error("cannot get metadata:%w", zap.Error(err))
		return nil
	}

	a.AppID = 0

	// if app id is greater than 0 then it must be created before submitting data
	if a.config.AppID != 0 {
		a.AppID = a.config.AppID
	}

	a.GenesisHash, err = a.API.RPC.Chain.GetBlockHash(0)
	if err != nil {
		// log.Error("cannot get genesis hash:%w", zap.Error(err))
		return nil
	}

	a.Rv, err = a.API.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		// log.Error("cannot get runtime version:%w", zap.Error(err))
		return nil
	}

	a.KeyringPair, err = signature.KeyringPairFromSecret(a.config.Seed, 42)
	if err != nil {
		// log.Error("cannot get keyring pair:%w", zap.Error(err))
		return nil
	}

	// DestinationAddress, err = types.NewHashFromHexString(Config.DestinationAddress)
	// if err != nil {
	// 	log.Fatalf("cannot decode destination address:%w", err)
	// }

	// DestinationDomain = types.NewUCompactFromUInt(uint64(Config.DestinationDomain))
	return &a
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
	batchDAData.Width = dataProof.NumberOfLeaves
	batchDAData.LeafIndex = dataProof.LeafIndex

	header, err := a.API.RPC.Chain.GetHeader(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get header", err)
	}

	batchDAData.BlockNumber = uint(header.Number)
	data, err := a.GetData(uint64(header.Number), dataProof.LeafIndex)
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
