package verify

import (
	"context"
	"encoding/hex"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/CryptoKass/blobstreamx-example/client"
	"github.com/celestiaorg/go-square/shares"
	"github.com/celestiaorg/go-square/square"
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

func (d *DAVerifier) VerifyDataAvailable(txHash string) (string, error) {
	shareRange, err := d.GetSharePointer(txHash)
	if err != nil {
		return "", errors.Wrap(err, "failed to get share range")
	}
	fmt.Println("Successfully got share range", shareRange)

	proof, root, err := client.GetShareProof(eth, trpc, sp)
	if err != nil {
		panic(fmt.Errorf("failed to get share proof: %w", err))
	}
	shares, err := shares.FromBytes(proof.Data)
	if err != nil {
		panic(fmt.Errorf("failed to parse shares: %w", err))
	}
	isCompactShare, err := shares[0].IsCompactShare()
	if err != nil {
		panic(fmt.Errorf("failed to check if share is compact: %w", err))
	}
	fmt.Println("isCompactShare", isCompactShare)

	fmt.Println("Share proof:", len(proof.Data), len(proof.RowProofs))
	fmt.Println("Blob data in share proof:", string(proof.Data[0][34:]))

	return "", nil
}

func (d *DAVerifier) GetSharePointer(txHash string) (square.ShareRange, error) {
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode transaction hash")
	}
	tx, err := d.tRPCClient.Tx(context.Background(), txHashBytes, true)
	if err != nil {
		return "", errors.Wrap(err, "failed to get transaction")
	}

	blockRes, err := d.tRPCClient.Block(context.Background(), &tx.Height)
	if err != nil {
		return "", errors.Wrap(err, "failed to get block")
	}

	// shareRange, err := square.BlobShareRange(blockRes.Block.Data.Txs.ToSliceOfBytes(), int(tx.Index), 0, blockRes.Block.Header.Version.App)
	shareRange, err := square.TxShareRange(blockRes.Block.Data.Txs.ToSliceOfBytes(), int(tx.Index), blockRes.Block.Header.Version.App)
	if err != nil {
		return "", errors.Wrap(err, "failed to get share range")
	}
	return &client.SharePointer{
		Height: tx.Height,
		Start:  int64(shareRange.Start),
		End:    int64(shareRange.End),
	}, nil
}
