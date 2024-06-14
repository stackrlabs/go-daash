package celestia

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/pkg/square"
	"github.com/celestiaorg/celestia-app/x/blob/types"
	rpc "github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/stackrlabs/go-daash/da"
)

// Client to interact with Celestia DA
type Client struct {
	lightClient *rpc.Client
	nodeClient  *http.HTTP
	Namespace   share.Namespace
	gasPrice    float64
	ctx         context.Context
}

// Returns an intialised Celestia DA client
func NewClient(ctx context.Context, lightClientRPCUrl string, nodeRPCUrl string, authToken string, hexNamespace string, gasPrice float64) (*Client, error) {
	nsBytes := make([]byte, 10)
	_, err := hex.Decode(nsBytes, []byte(hexNamespace))
	if err != nil {
		log.Fatalln("invalid hex value of a namespace:", err)
		return nil, err
	}
	namespace, err := share.NewBlobNamespaceV0(nsBytes)
	if err != nil {
		return nil, err
	}
	client, err := rpc.NewClient(ctx, lightClientRPCUrl, authToken)
	if err != nil {
		fmt.Printf("failed to create rpc client: %v", err)
		return nil, err
	}
	nodeClient, err := http.New(nodeRPCUrl, "/websocket")
	if err != nil {
		return nil, err
	}
	return &Client{
		lightClient: client,
		nodeClient:  nodeClient,
		Namespace:   namespace,
		gasPrice:    gasPrice,
		ctx:         ctx,
	}, nil
}

// MaxBlobSize returns the max blob size
func (c *Client) MaxBlobSize(ctx context.Context) (uint64, error) {
	// TODO: pass-through query to node, app
	return appconsts.DefaultMaxBytes, nil
}

// Get returns Blob for each given ID, or an error.
func (c *Client) Get(ctx context.Context, id da.ID) (da.Blob, error) {
	celestiaID, ok := id.(ID)
	if !ok {
		return nil, errors.New("invalid ID")
	}
	blob, err := c.lightClient.Blob.Get(ctx, celestiaID.Height, c.Namespace, celestiaID.ShareCommitment)
	if err != nil {
		return nil, err
	}

	return blob.Data, nil
}

// Commit creates a Commitment for each given Blob.
func (c *Client) Commit(ctx context.Context, daBlob da.Blob) (da.Commitment, error) {
	blob, err := blob.NewBlobV0(c.Namespace, daBlob)
	if err != nil {
		return nil, err
	}
	commitment, err := types.CreateCommitment(&blob.Blob)
	if err != nil {
		return nil, err
	}
	return commitment, err
}

// Submit submits the Blobs to Data Availability layer.
func (c *Client) Submit(ctx context.Context, daBlob da.Blob, gasPrice float64) (da.ID, error) {
	b, err := blob.NewBlobV0(c.Namespace, daBlob)
	if err != nil {
		return nil, err
	}
	options := blob.DefaultSubmitOptions()
	// if gas price was configured globally use that as the default
	if c.gasPrice >= 0 && gasPrice < 0 {
		gasPrice = c.gasPrice
	}
	if gasPrice >= 0 {
		options.GasLimit = types.EstimateGas([]uint32{uint32(len(b.Data))}, appconsts.DefaultGasPerBlobByte, auth.DefaultTxSizeCostPerByte)
		options.Fee = sdktypes.NewInt(int64(math.Ceil(gasPrice * float64(options.GasLimit)))).Int64()
	}
	txResp, err := c.lightClient.State.SubmitPayForBlob(ctx, sdktypes.NewInt(int64(options.Fee)), options.GasLimit, []*blob.Blob{b})
	if err != nil {
		return nil, err
	}

	log.Println("successfully submitted blobs", "height", txResp.Height, "gas", options.GasLimit, "fee", options.Fee)

	commitment, err := c.Commit(ctx, daBlob)
	if err != nil {
		return nil, err
	}
	shareCommitment, ok := commitment.(blob.Commitment)
	if !ok {
		return nil, errors.New("invalid commitment")
	}
	sp, err := c.getSharePointer(ctx, txResp.TxHash)
	if err != nil {
		return nil, err
	}
	id := ID{
		Height:          uint64(txResp.Height),
		ShareCommitment: shareCommitment,
		TxHash:          txResp.TxHash,
		SharePointer:    sp,
	}

	return id, nil
}

func (c *Client) GetProof(ctx context.Context, id da.ID) (da.Proof, error) {
	celestiaID, ok := id.(ID)
	if !ok {
		return nil, errors.New("invalid ID")
	}
	proof, err := c.lightClient.Blob.GetProof(ctx, celestiaID.Height, c.Namespace, celestiaID.ShareCommitment)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

// Validate validates Commitments against the corresponding Proofs. This should be possible without retrieving the Blobs.
func (c *Client) Validate(ctx context.Context, id da.ID, daProof da.Proof) (bool, error) {
	celestiaID, ok := id.(ID)
	if !ok {
		return false, errors.New("invalid ID")
	}
	proof, ok := daProof.(Proof)
	if !ok {
		return false, errors.New("invalid proof")
	}
	// TODO(tzdybal): for some reason, if proof doesn't match commitment, API returns (false, "blob: invalid proof")
	//    but analysis of the code in celestia-node implies this should never happen - maybe it's caused by openrpc?
	//    there is no way of gently handling errors here, but returned value is fine for us
	isIncluded, _ := c.lightClient.Blob.Included(ctx, celestiaID.Height, c.Namespace, &proof, celestiaID.ShareCommitment)
	return isIncluded, nil
}

// getSharePointer returns the share pointer for the given transaction hash.
func (c *Client) getSharePointer(ctx context.Context, txHash string) (SharePointer, error) {
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to decode transaction hash: %w", err)
	}
	tx, err := c.nodeClient.Tx(ctx, txHashBytes, true)
	if err != nil {
		return SharePointer{}, fmt.Errorf("failed to get transaction: %w", err)
	}

	blockRes, err := c.nodeClient.Block(ctx, &tx.Height)
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
