package verify

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"errors"

	"github.com/celestiaorg/celestia-app/pkg/shares"
	"github.com/celestiaorg/celestia-app/pkg/square"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	bv "github.com/stackrlabs/go-daash/celestia/verify/bindings/blobstreamverifier"
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

func (d *Verifier) VerifyDataAvailable(txHash string) (bool, error) {
	shareRange, err := d.GetSharePointer(txHash)
	if err != nil {
		return false, fmt.Errorf("failed to get share range: %w", err)
	}
	fmt.Println("Successfully got share range", shareRange)

	proof, root, err := GetShareProof(d.ethClient, d.tRPCClient, &shareRange, d.blobstreamXContract)
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
			Height: big.NewInt(shareRange.Height),
			Index:  big.NewInt(shareRange.Start),
			Length: big.NewInt(shareRange.End - shareRange.Start),
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

func (d *Verifier) GetSharePointer(txHash string) (SharePointer, error) {
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to decode transaction hash: %w", err)
	}
	tx, err := d.tRPCClient.Tx(context.Background(), txHashBytes, true)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to get transaction: %w", err)
	}

	blockRes, err := d.tRPCClient.Block(context.Background(), &tx.Height)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to get block: %w", err)
	}

	shareRange, err := square.BlobShareRange(blockRes.Block.Data.Txs.ToSliceOfBytes(), int(tx.Index), 0, blockRes.Block.Header.Version.App)
	// shareRange, err := square.TxShareRange(blockRes.Block.Data.Txs.ToSliceOfBytes(), int(tx.Index), blockRes.Block.Header.Version.App)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to get share range: %w", err)
	}
	return SharePointer{
		Height: tx.Height,
		Start:  int64(shareRange.Start),
		End:    int64(shareRange.End),
	}, nil
}
