package verify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/availda/verify/bindings/vectorverifier"
)

const (
	succinctBaseURL = "https://beaconapi.succinct.xyz/api/integrations/vectorx"
)

// Verfifier is used to verify availability of Avail blobs on EVM chains
type Verifier struct {
	daClient         *availda.DAClient
	ethClient        *ethclient.Client
	vectorXContract  common.Address
	bridgeContract   common.Address
	verifierContract common.Address
	availNetwork     string
}

type SuccinctAPIResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    struct {
		BlockHash      string   `json:"blockHash"`
		DataCommitment string   `json:"dataCommitment"`
		DataRoot       string   `json:"dataRoot"`
		Index          int64    `json:"index"`
		MerkleBranch   []string `json:"merkleBranch"`
		RangeHash      string   `json:"rangeHash"`
		TotalLeaves    int64    `json:"totalLeaves"`
	} `json:"data"`
}

func NewVerifier(client *availda.DAClient, ethEndpoint string, bridgeContract string, verifierContract string, vectorXContract string, availNetwork string) (*Verifier, error) {
	ethClient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create eth client: %w", err)
	}
	return &Verifier{
		daClient:         client,
		ethClient:        ethClient,
		bridgeContract:   common.HexToAddress(bridgeContract),
		verifierContract: common.HexToAddress(verifierContract),
		vectorXContract:  common.HexToAddress(vectorXContract),
		availNetwork:     availNetwork,
	}, nil
}

func (d *Verifier) IsDataAvailable(blockHeight uint64, extIndex uint64) (bool, error) {
	dataProof, err := d.daClient.GetProof(context.Background(), uint32(blockHeight), int(extIndex))
	if err != nil {
		return false, fmt.Errorf("failed to get data proof: %w", err)
	}
	proof, err := d.GetAggregatedProof(dataProof, blockHeight)
	if err != nil {
		return false, fmt.Errorf("failed to get aggregated proof: %w", err)
	}
	verifier, err := vectorverifier.NewVectorverifier(d.verifierContract, d.ethClient)
	if err != nil {
		return false, fmt.Errorf("failed to create vector verifier: %w", err)
	}
	success, err := verifier.VerifyDataAvailability(nil, d.bridgeContract, proof.Leaf, *proof)
	if err != nil {
		return false, fmt.Errorf("failed to verify blob leaf: %w", err)
	}
	return success, nil
}

// IsDataIncluded verifies that the blob data corresponding to the given block height and external index is available on DA
func (d *Verifier) IsDataIncluded(blockHeight uint64, extIndex uint64) (bool, error) {
	id := availda.MakeID(uint32(blockHeight), int(extIndex))
	blobs, err := d.daClient.Get(context.Background(), []da.ID{id})
	if err != nil {
		return false, fmt.Errorf("failed to get blob data: %w", err)
	}
	fmt.Println("size of blob data:", len(blobs[0]))

	dataProof, err := d.daClient.GetProof(context.Background(), uint32(blockHeight), int(extIndex))
	if err != nil {
		return false, fmt.Errorf("failed to get data proof: %w", err)
	}
	proof, err := d.GetAggregatedProof(dataProof, blockHeight)
	if err != nil {
		return false, fmt.Errorf("failed to get aggregated proof: %w", err)
	}
	verifier, err := vectorverifier.NewVectorverifier(d.verifierContract, d.ethClient)
	if err != nil {
		return false, fmt.Errorf("failed to create vector verifier: %w", err)
	}
	success, err := verifier.VerifyDataInclusion(nil, d.bridgeContract, blobs[0], *proof)
	if err != nil {
		return false, fmt.Errorf("failed to verify blob leaf: %w", err)
	}
	return success, nil
}

func (d *Verifier) GetAggregatedProof(dataProof availda.DataProofRPCResponse, blockHeight uint64) (*vectorverifier.IAvailBridgeMerkleProofInput, error) {
	chainID, err := d.ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cannot get chain id:%w", err)
	}
	blockHash, err := d.daClient.API.RPC.Chain.GetBlockHash(blockHeight)
	if err != nil {
		return nil, fmt.Errorf("cannot get block hash:%w", err)
	}
	resp, err := http.Get(
		fmt.Sprintf("%s?chainName=%s&contractChainId=%s&contractAddress=%s&blockHash=%s",
			succinctBaseURL,
			d.availNetwork,
			chainID.String(),
			d.vectorXContract.Hex(),
			blockHash.Hex(),
		))
	if err != nil {
		return nil, fmt.Errorf("cannot get succinct proof:%w", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read data:%w", err)
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot close body:%w", err)
	}
	fmt.Println("raw succinct response", string(data))
	var succinctAPIResponse SuccinctAPIResponse
	err = json.Unmarshal(data, &succinctAPIResponse)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal data:%w", err)
	}
	if succinctAPIResponse.Error != "" {
		return nil, fmt.Errorf("succinct api failed: %s", succinctAPIResponse.Error)
	}

	var dataRootProof [][32]byte
	for _, node := range succinctAPIResponse.Data.MerkleBranch {
		hexBytes := common.HexToHash(node)
		dataRootProof = append(dataRootProof, hexBytes)
	}
	var leafProof [][32]byte
	for _, node := range dataProof.Result.DataProof.Proof {
		hexBytes := common.HexToHash(node)
		leafProof = append(leafProof, hexBytes)
	}

	return &vectorverifier.IAvailBridgeMerkleProofInput{
		DataRootProof: dataRootProof,
		LeafProof:     leafProof,
		RangeHash:     common.HexToHash(succinctAPIResponse.Data.RangeHash),
		DataRootIndex: big.NewInt(succinctAPIResponse.Data.Index),
		BlobRoot:      common.HexToHash(dataProof.Result.DataProof.Roots.BlobRoot),
		BridgeRoot:    common.HexToHash(dataProof.Result.DataProof.Roots.BridgeRoot),
		Leaf:          common.HexToHash(dataProof.Result.DataProof.Leaf),
		LeafIndex:     big.NewInt(dataProof.Result.DataProof.LeafIndex),
	}, nil
}
