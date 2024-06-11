package verify

import (
	"fmt"
	"math/big"

	"errors"

	"github.com/celestiaorg/celestia-app/pkg/shares"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stackrlabs/go-daash/celestia"
	bv "github.com/stackrlabs/go-daash/celestia/verify/blobstream/verifier"
	"github.com/tendermint/tendermint/rpc/client/http"
)

type Verifier struct {
	ethClient           *ethclient.Client
	tRPCClient          *http.HTTP
	verifierContract    common.Address
	blobstreamXContract common.Address
}

func NewVerifier(ethEndpoint string, tRPCEndpoint string, verifierContract string, blobstreamXContract string) (*Verifier, error) {
	ethClient, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, err
	}
	trpc, err := http.New(tRPCEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}
	return &Verifier{
		ethClient:           ethClient,
		tRPCClient:          trpc,
		verifierContract:    common.HexToAddress(verifierContract),
		blobstreamXContract: common.HexToAddress(blobstreamXContract),
	}, nil
}

func (d *Verifier) VerifyDataAvailable(id celestia.ID) (bool, error) {
	sp := id.SharePointer
	proof, root, err := d.GetShareProof(sp)
	if err != nil {
		return false, fmt.Errorf("failed to get share proof: %w", err)
	}
	shares, err := shares.FromBytes(proof.Data)
	if err != nil {
		return false, fmt.Errorf("failed to parse shares: %w", err)
	}
	isCompactShare, err := shares[0].IsCompactShare()
	if err != nil {
		return false, fmt.Errorf("failed to check if share is compact: %w", err)
	}
	fmt.Println("isCompactShare", isCompactShare)

	fmt.Println("Share proof:", len(proof.Data), len(proof.RowProofs))
	fmt.Println("Blob data in share proof:", string(proof.Data[0][34:]))

	verifier, err := bv.NewBlobstreamverifier(d.verifierContract, d.ethClient)
	if err != nil {
		return false, fmt.Errorf("failed to create new blobstream verifier: %w", err)
	}

	success, err := verifier.VerifyDataAvailability(
		nil,
		d.blobstreamXContract,
		bv.SpanSequence{
			Height: big.NewInt(sp.Height),
			Index:  big.NewInt(sp.Start),
			Length: big.NewInt(sp.End - sp.Start),
		},
		proof.RowRoots,
		proof.RowProofs,
		proof.AttestationProof,
		root,
	)
	if err != nil {
		return false, fmt.Errorf("failed to verify data availability: %w", err)
	}
	if !success {
		return false, errors.New("failed to verify data availability")
	}

	return true, nil
}
