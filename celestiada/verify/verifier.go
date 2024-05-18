package verify

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/rpc/client/http"
)

type DAVerifier struct {
	ethClient           *ethclient.Client
	tRPCClient          *http.HTTP
	verifierContract    common.Address
	blobstreamXContract common.Address
}

func NewDAVerifier(ethEndpoint string, tRPCEndpoint string, verifierContract common.Address, blobstreamXContract common.Address) (*DAVerifier, error) {
	ethClient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, err
	}
	trpc, err := http.New(tRPCEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}
	return &DAVerifier{
		ethClient:           ethClient,
		tRPCClient:          trpc,
		verifierContract:    verifierContract,
		blobstreamXContract: blobstreamXContract,
	}, nil
}
