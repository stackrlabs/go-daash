package verify

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rollkit/go-da"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/availda/verify/bindings/vectorverifier"
)

type DAVerifier struct {
	daClient         *availda.DAClient
	ethClient        *ethclient.Client
	bridgeContract   common.Address
	verifierContract common.Address
}

func NewDAVerifier(configPath string, ethEndpoint string, bridgeContract string, verifierContract string) (*DAVerifier, error) {
	client, err := availda.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create DAClient: %w", err)
	}
	ethClient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create eth client: %w", err)
	}
	return &DAVerifier{
		daClient:         client,
		ethClient:        ethClient,
		bridgeContract:   common.HexToAddress(bridgeContract),
		verifierContract: common.HexToAddress(verifierContract),
	}, nil
}

func (d *DAVerifier) VerifyDataAvailable(blockHeight uint64, extIndex uint64) (bool, error) {
	proof, err := d.daClient.GetProof(context.Background(), uint32(blockHeight), int(extIndex))
	if err != nil {
		return false, fmt.Errorf("failed to get proof: %w", err)
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

// VerifyDataIncluded verifies that the blob data corresponding to the given block height and external index is available on DA
func (d *DAVerifier) VerifyDataIncluded(blockHeight uint64, extIndex uint64) (bool, error) {
	id := availda.MakeID(uint32(blockHeight), int(extIndex))
	blobs, err := d.daClient.Get(context.Background(), []da.ID{id})
	if err != nil {
		return false, fmt.Errorf("failed to get blob data: %w", err)
	}
	fmt.Println("size of blob data:", len(blobs[0]))

	proof, err := d.daClient.GetProof(context.Background(), uint32(blockHeight), int(extIndex))
	if err != nil {
		return false, fmt.Errorf("failed to get proof: %w", err)
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
