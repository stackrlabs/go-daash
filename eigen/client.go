package eigen

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Layr-Labs/eigenda/api/clients"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stackrlabs/go-daash/da"
)

type Client struct {
	// DisperserClient is the gRPC client for the Disperser service.
	client clients.EigenDAClient
}

// NewClient returns a new instance of the EigenDA client.
func NewClient(config clients.EigenDAClientConfig) (*Client, error) {
	logger := log.NewLogger(slog.NewTextHandler(os.Stdout, nil))
	client, err := clients.NewEigenDAClient(logger, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create EigenDA client: %v", err)
	}

	return &Client{
		client: *client,
	}, nil
}

func (e *Client) MaxBlobSize(ctx context.Context) (uint64, error) {
	return 512 * 1024, nil // Currently set at 512KB
}

func (c *Client) Submit(ctx context.Context, daBlob da.Blob, gasPrice float64) (da.ID, error) {
	blobInfo, err := c.client.PutBlob(ctx, daBlob)
	if err != nil {
		return nil, fmt.Errorf("failed to disperse blob: %v", err)
	}
	blobID := ID{
		BlobIndex:       blobInfo.BlobVerificationProof.BlobIndex,
		BatchHeaderHash: blobInfo.BlobVerificationProof.BatchMetadata.BatchHeaderHash,
	}
	return blobID, nil
}

func (c *Client) Get(ctx context.Context, id da.ID) (da.Blob, error) {
	blobID, ok := id.(ID)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}
	blob, err := c.client.GetBlob(ctx, blobID.BatchHeaderHash, blobID.BlobIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve blob: %v", err)
	}

	return blob, nil
}

func (e *Client) Commit(ctx context.Context, daBlob da.Blob) (da.Commitment, error) {
	return nil, nil
}

func (e *Client) Validate(ctx context.Context, id da.ID, proof da.Proof) (bool, error) {
	return false, nil
}

func (e *Client) GetProof(ctx context.Context, id da.ID) (da.Proof, error) {
	return nil, nil
}
