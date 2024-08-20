package eigen

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Layr-Labs/eigenda/api/clients"
	grpcdisperser "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/Layr-Labs/eigenda/disperser"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stackrlabs/go-daash/da"
)

type Client struct {
	// internalClient is used to interact with the EigenDA API
	internalClient clients.EigenDAClient
}

// NewClient returns a new instance of the EigenDA client.
func NewClient(config clients.EigenDAClientConfig) (*Client, error) {
	logger := log.NewLogger(slog.NewTextHandler(os.Stdout, nil))
	client, err := clients.NewEigenDAClient(logger, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create EigenDA client: %v", err)
	}

	return &Client{
		internalClient: *client,
	}, nil
}

func (e *Client) MaxBlobSize(ctx context.Context) (uint64, error) {
	return 2 * 1024 * 1024, nil // Currently set at 2MB
}

func (c *Client) Submit(ctx context.Context, daBlob da.Blob, gasPrice float64) (da.ID, error) {
	start := time.Now()
	blobID, err := c.PutBlob(ctx, daBlob)
	if err != nil {
		return nil, fmt.Errorf("failed to disperse blob: %v", err)
	}
	end := time.Now()
	fmt.Println("Time taken to disperse blob:", end.Sub(start))

	return blobID, nil
}

func (c *Client) Get(ctx context.Context, id da.ID) (da.Blob, error) {
	blobID, ok := id.(*ID)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}
	blob, err := c.internalClient.GetBlob(ctx, blobID.BlobInfo.BlobVerificationProof.BatchMetadata.BatchHeaderHash, blobID.BlobInfo.BlobVerificationProof.BlobIndex)
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

// PutBlob encodes and writes a blob to EigenDA, waiting for it to be confirmed
// before returning. This function is resiliant to transient failures and
// timeouts.
func (c *Client) PutBlob(ctx context.Context, data []byte) (ID, error) {
	resultChan, errorChan := c.PutBlobAsync(ctx, data)
	select { // no timeout here because we depend on the configured timeout in PutBlobAsync
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return ID{}, err
	}
}

func (c *Client) PutBlobAsync(ctx context.Context, data []byte) (resultChan chan ID, errChan chan error) {
	resultChan = make(chan ID, 1)
	errChan = make(chan error, 1)
	go c.putBlob(ctx, data, resultChan, errChan)
	return
}

func (c *Client) putBlob(ctx context.Context, rawData []byte, resultChan chan ID, errChan chan error) {
	// encode blob
	if c.internalClient.Codec == nil {
		errChan <- fmt.Errorf("codec cannot be nil")
		return
	}

	data, err := c.internalClient.Codec.EncodeBlob(rawData)
	if err != nil {
		errChan <- fmt.Errorf("error encoding blob: %w", err)
		return
	}

	customQuorumNumbers := make([]uint8, len(c.internalClient.Config.CustomQuorumIDs))
	for i, id := range c.internalClient.Config.CustomQuorumIDs {
		customQuorumNumbers[i] = uint8(id)
	}
	// disperse blob
	blobStatus, requestID, err := c.internalClient.Client.DisperseBlobAuthenticated(ctx, data, customQuorumNumbers)
	if err != nil {
		errChan <- fmt.Errorf("error initializing DisperseBlobAuthenticated() client: %w", err)
		return
	}

	// process response
	if *blobStatus == disperser.Failed {
		errChan <- fmt.Errorf("reply status is %d", blobStatus)
		return
	}

	base64RequestID := base64.StdEncoding.EncodeToString(requestID)
	fmt.Println("Blob dispersed to EigenDA, now waiting for confirmation", "requestID", base64RequestID)

	ticker := time.NewTicker(c.internalClient.Config.StatusQueryRetryInterval)
	defer ticker.Stop()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, c.internalClient.Config.StatusQueryTimeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			errChan <- fmt.Errorf("timed out waiting for EigenDA blob to confirm blob with request id=%s: %w", base64RequestID, ctx.Err())
			return
		case <-ticker.C:
			statusRes, err := c.internalClient.Client.GetBlobStatus(ctx, requestID)
			if err != nil {
				c.internalClient.Log.Error("Unable to retrieve blob dispersal status, will retry", "requestID", base64RequestID, "err", err)
				continue
			}

			switch statusRes.Status {
			case grpcdisperser.BlobStatus_PROCESSING, grpcdisperser.BlobStatus_DISPERSING:
				fmt.Println("Blob submitted, waiting for dispersal from EigenDA", "requestID", base64RequestID)
			case grpcdisperser.BlobStatus_FAILED:
				errChan <- fmt.Errorf("EigenDA blob dispersal failed in processing, requestID=%s: %w", base64RequestID, err)
				return
			case grpcdisperser.BlobStatus_INSUFFICIENT_SIGNATURES:
				errChan <- fmt.Errorf("EigenDA blob dispersal failed in processing with insufficient signatures, requestID=%s: %w", base64RequestID, err)
				return
			case grpcdisperser.BlobStatus_CONFIRMED:
				fmt.Println("EigenDA blob confirmed, waiting for finalization", "requestID", base64RequestID)
				resultChan <- ID{BlobInfo: statusRes.Info, RequestID: string(requestID)}
			case grpcdisperser.BlobStatus_FINALIZED:
				batchHeaderHashHex := fmt.Sprintf("0x%s", hex.EncodeToString(statusRes.Info.BlobVerificationProof.BatchMetadata.BatchHeaderHash))
				fmt.Println("Successfully dispersed blob to EigenDA", "requestID", base64RequestID, "batchHeaderHash", batchHeaderHashHex)
				return
			default:
				errChan <- fmt.Errorf("EigenDA blob dispersal failed in processing with reply status %d", statusRes.Status)
				return
			}
		}
	}
}
