package verify

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stackrlabs/go-daash/availda"
	"github.com/stackrlabs/go-daash/availda/verify/bindings/availbridge"
)

type DAVerifier struct {
	daClient       *availda.DAClient
	ethClient      *ethclient.Client
	bridgeContract common.Address
}

func NewDAVerifier(configPath string, ethEndpoint string, bridgeContract string) (*DAVerifier, error) {
	client, err := availda.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create DAClient: %w", err)
	}
	ethClient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create eth client: %w", err)
	}
	return &DAVerifier{
		daClient:       client,
		ethClient:      ethClient,
		bridgeContract: common.HexToAddress(bridgeContract),
	}, nil
}

func (d *DAVerifier) VerifyDataAvailable(blockHeight uint64, extIndex uint64) (bool, error) {
	proof, err := d.daClient.GetProof(context.Background(), uint32(blockHeight), int(extIndex))
	if err != nil {
		return false, fmt.Errorf("failed to get proof: %w", err)
	}
	bridge, err := availbridge.NewAvailbridge(d.bridgeContract, d.ethClient)
	if err != nil {
		return false, fmt.Errorf("failed to create availbridge: %w", err)
	}
	success, err := bridge.VerifyBlobLeaf(nil, *proof)
	if err != nil {
		return false, fmt.Errorf("failed to verify blob leaf: %w", err)
	}
	return success, nil
}
