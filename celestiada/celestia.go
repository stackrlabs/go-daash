package celestiada

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/x/blob/types"
	rpc "github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/nmt"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/rollkit/go-da"
)

// CelestiaDA implements the celestia backend for the DA interface
type DAClient struct {
	client    *rpc.Client
	Namespace share.Namespace
	gasPrice  float64
	ctx       context.Context
}

// Returns an intialised Celestia DA client
func New(ctx context.Context, lightCLientRPCUrl string, authToken string, hexNamespace string, gasPrice float64) (*DAClient, error) {
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
	client, err := rpc.NewClient(ctx, lightCLientRPCUrl, authToken)
	if err != nil {
		fmt.Printf("failed to create rpc client: %v", err)
		return nil, err
	}
	return &DAClient{
		client:    client,
		Namespace: namespace,
		gasPrice:  gasPrice,
		ctx:       ctx,
	}, nil
}

// MaxBlobSize returns the max blob size
func (c *DAClient) MaxBlobSize(ctx context.Context) (uint64, error) {
	// TODO: pass-through query to node, app
	return appconsts.DefaultMaxBytes, nil
}

// Get returns Blob for each given ID, or an error.
func (c *DAClient) Get(ctx context.Context, ids []da.ID) ([]da.Blob, error) {
	var blobs []da.Blob
	for _, id := range ids {
		height, _, commitment := SplitID(id)
		blob, err := c.client.Blob.Get(ctx, height, c.Namespace, commitment)
		if err != nil {
			return nil, err
		}
		blobs = append(blobs, blob.Data)
	}
	return blobs, nil
}

// GetIDs returns IDs of all Blobs located in DA at given height.
func (c *DAClient) GetIDs(ctx context.Context, height uint64) ([]da.ID, error) {
	var ids []da.ID
	blobs, err := c.client.Blob.GetAll(ctx, height, []share.Namespace{c.Namespace})
	if err != nil {
		if strings.Contains(err.Error(), blob.ErrBlobNotFound.Error()) {
			return nil, nil
		}
		return nil, err
	}
	for _, b := range blobs {
		ids = append(ids, makeID(height, b.Commitment, make([]byte, 32)))
	}
	return ids, nil
}

// Commit creates a Commitment for each given Blob.
func (c *DAClient) Commit(ctx context.Context, daBlobs []da.Blob) ([]da.Commitment, error) {
	_, commitments, err := c.blobsAndCommitments(daBlobs)
	return commitments, err
}

// Submit submits the Blobs to Data Availability layer.
func (c *DAClient) Submit(ctx context.Context, daBlobs []da.Blob, gasPrice float64) ([]da.ID, []da.Proof, error) {
	blobs, commitments, err := c.blobsAndCommitments(daBlobs)
	if err != nil {
		return nil, nil, err
	}
	options := blob.DefaultSubmitOptions()
	// if gas price was configured globally use that as the default
	if c.gasPrice >= 0 && gasPrice < 0 {
		gasPrice = c.gasPrice
	}
	if gasPrice >= 0 {
		blobSizes := make([]uint32, len(blobs))
		for i, blob := range blobs {
			blobSizes[i] = uint32(len(blob.Data))
		}
		options.GasLimit = types.EstimateGas(blobSizes, appconsts.DefaultGasPerBlobByte, auth.DefaultTxSizeCostPerByte)
		options.Fee = sdktypes.NewInt(int64(math.Ceil(gasPrice * float64(options.GasLimit)))).Int64()
	}
	txResp, err := c.client.State.SubmitPayForBlob(ctx, sdktypes.NewInt(int64(options.Fee)), options.GasLimit, blobs)
	if err != nil {
		return nil, nil, err
	}

	log.Println("successfully submitted blobs", "height", txResp.Height, "gas", options.GasLimit, "fee", options.Fee)
	ids := make([]da.ID, len(daBlobs))
	proofs := make([]da.Proof, len(daBlobs))
	for i, commitment := range commitments {
		txHashBytes, err := hex.DecodeString(txResp.TxHash)
		if err != nil {
			return nil, nil, err
		}
		ids[i] = makeID(uint64(txResp.Height), commitment, txHashBytes)
		proof, err := c.client.Blob.GetProof(ctx, uint64(txResp.Height), c.Namespace, commitment)
		if err != nil {
			return nil, nil, err
		}
		// TODO(tzdybal): does always len(*proof) == 1?
		proofs[i], err = (*proof)[0].MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
	}
	return ids, proofs, nil
}

// blobsAndCommitments converts []da.Blob to []*blob.Blob and generates corresponding []da.Commitment
func (c *DAClient) blobsAndCommitments(daBlobs []da.Blob) ([]*blob.Blob, []da.Commitment, error) {
	var blobs []*blob.Blob
	var commitments []da.Commitment
	for _, daBlob := range daBlobs {
		b, err := blob.NewBlobV0(c.Namespace, daBlob)
		if err != nil {
			return nil, nil, err
		}
		blobs = append(blobs, b)

		commitment, err := types.CreateCommitment(&b.Blob)
		if err != nil {
			return nil, nil, err
		}
		commitments = append(commitments, commitment)
	}
	return blobs, commitments, nil
}

// Validate validates Commitments against the corresponding Proofs. This should be possible without retrieving the Blobs.
func (c *DAClient) Validate(ctx context.Context, ids []da.ID, daProofs []da.Proof) ([]bool, error) {
	var included []bool
	var proofs []*blob.Proof
	for _, daProof := range daProofs {
		nmtProof := &nmt.Proof{}
		if err := nmtProof.UnmarshalJSON(daProof); err != nil {
			return nil, err
		}
		proof := &blob.Proof{nmtProof}
		proofs = append(proofs, proof)
	}
	for i, id := range ids {
		height, _, commitment := SplitID(id)
		// TODO(tzdybal): for some reason, if proof doesn't match commitment, API returns (false, "blob: invalid proof")
		//    but analysis of the code in celestia-node implies this should never happen - maybe it's caused by openrpc?
		//    there is no way of gently handling errors here, but returned value is fine for us
		isIncluded, _ := c.client.Blob.Included(ctx, height, c.Namespace, proofs[i], commitment)
		included = append(included, isIncluded)
	}
	return included, nil
}

// heightLen is a length (in bytes) of serialized height.
//
// This is 8 as uint64 consist of 8 bytes.
const heightLen = 8
const txHashLen = 32

func makeID(height uint64, commitment da.Commitment, txHash []byte) da.ID {
	id := make([]byte, heightLen+txHashLen+len(commitment))
	binary.LittleEndian.PutUint64(id, height)
	copy(id[heightLen:heightLen+txHashLen], txHash)
	copy(id[heightLen+txHashLen:], commitment)
	return id
}

func SplitID(id da.ID) (uint64, []byte, da.Commitment) {
	if len(id) <= heightLen {
		return 0, nil, nil
	}
	// Return only height and commitment if ID is smaller to keep it backward compatible
	if len(id) <= heightLen+txHashLen {
		return 0, nil, id[heightLen : heightLen+txHashLen]
	}
	return binary.LittleEndian.Uint64(id[:heightLen]), id[heightLen : heightLen+txHashLen], id[heightLen+txHashLen:]
}
