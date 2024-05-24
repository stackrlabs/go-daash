package eigen

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/stackrlabs/go-daash/da"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	// DaRpc is the HTTP provider URL for the Data Availability node.
	DARpc string

	// DisperserClient is the gRPC client for the Disperser service.
	disperserClient DisperserClient

	// Quorum IDs and SecurityParams to use when dispersing and retrieving blobs
	DADisperserSecurityParams []*SecurityParams

	// The total amount of time that the batcher will spend waiting for EigenDA to confirm a blob
	DAStatusQueryTimeout time.Duration

	// The amount of time to wait between status queries of a newly dispersed blob
	DAStatusQueryRetryInterval time.Duration
}

// NewClient returns a new instance of the EigenDA client.
func NewClient(daRpc string, daStatusQueryTimeout time.Duration, daStatusQueryRetryInterval time.Duration) (*Client, error) {
	conn, err := grpc.Dial(daRpc, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		fmt.Println("Unable to connect to EigenDA, aborting", "err", err)
		return nil, err
	}
	daClient := NewDisperserClient(conn)

	disperserSecurityParams := []*SecurityParams{}
	disperserSecurityParams = append(disperserSecurityParams, &SecurityParams{
		QuorumId:           0,
		AdversaryThreshold: 25,
		QuorumThreshold:    50,
	})
	log.Println("🟢 EigenDA client initalised")
	return &Client{
		DARpc:                      daRpc,
		disperserClient:            daClient,
		DADisperserSecurityParams:  disperserSecurityParams,
		DAStatusQueryTimeout:       daStatusQueryTimeout,
		DAStatusQueryRetryInterval: daStatusQueryRetryInterval,
	}, nil
}

func (e *Client) MaxBlobSize(ctx context.Context) (uint64, error) {
	return 512 * 1024, nil // Currently set at 512KB
}

func (e *Client) Submit(ctx context.Context, daBlobs []da.Blob, gasPrice float64) ([]da.ID, []da.Proof, error) {
	blobInfo, err := e.disperseBlob(ctx, daBlobs[0])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to disperse blob: %v", err)
	}
	blobID := ID{
		BlobIndex:       blobInfo.BlobVerificationProof.BlobIndex,
		BatchHeaderHash: blobInfo.BlobVerificationProof.BatchMetadata.BatchHeaderHash,
	}
	return []da.ID{blobID}, []da.Proof{blobInfo.BlobVerificationProof.InclusionProof}, nil
}

func (e *Client) Get(ctx context.Context, ids []da.ID) ([]da.Blob, error) {
	blobID, ok := ids[0].(ID)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}
	resp, err := e.disperserClient.RetrieveBlob(ctx, &RetrieveBlobRequest{
		BlobIndex:       blobID.BlobIndex,
		BatchHeaderHash: blobID.BatchHeaderHash,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve blob: %v", err)
	}
	return []da.Blob{resp.Data}, nil
}

func (e *Client) GetIDs(ctx context.Context, height uint64) ([]da.ID, error) {
	return nil, nil
}

func (e *Client) Commit(ctx context.Context, daBlobs []da.Blob) ([]da.Commitment, error) {
	return nil, nil
}

func (e *Client) Validate(ctx context.Context, ids []da.ID, proofs []da.Proof) ([]bool, error) {
	return nil, nil
}

type ID struct {
	BlobIndex       uint32
	BatchHeaderHash []byte
}

func (e *Client) disperseBlob(ctx context.Context, txData []byte) (*BlobInfo, error) {
	fmt.Println("Attempting to disperse blob to EigenDA")

	disperseReq := &DisperseBlobRequest{
		Data:           txData,
		SecurityParams: e.DADisperserSecurityParams,
	}
	daClient := e.disperserClient
	disperseRes, err := daClient.DisperseBlob(ctx, disperseReq)
	fmt.Println("DisperseBlob response", "disperseRes", disperseRes, "err", err)

	if err != nil {
		fmt.Printf("Unable to disperse blob to EigenDA, aborting", "err", err)
		return nil, err
	}

	if disperseRes.Result == BlobStatus_UNKNOWN ||
		disperseRes.Result == BlobStatus_FAILED {
		fmt.Printf("Unable to disperse blob to EigenDA, aborting", "err", err)
		return nil, fmt.Errorf("reply status is %d", disperseRes.Result)
	}

	base64RequestID := base64.StdEncoding.EncodeToString(disperseRes.RequestId)

	fmt.Println("Blob disepersed to EigenDA, now waiting for confirmation", "requestID", base64RequestID)

	var statusRes *BlobStatusReply
	timeoutTime := time.Now().Add(e.DAStatusQueryTimeout)
	// Wait before first status check
	time.Sleep(e.DAStatusQueryRetryInterval)
	for time.Now().Before(timeoutTime) {
		statusRes, err = daClient.GetBlobStatus(ctx, &BlobStatusRequest{
			RequestId: disperseRes.RequestId,
		})
		if err != nil {
			fmt.Printf("Unable to retrieve blob dispersal status, will retry", "requestID", base64RequestID, "err", err)
		} else if statusRes.Status == BlobStatus_CONFIRMED {
			// TODO(eigenlayer): As long as fault proofs are disabled, we can move on once a blob is confirmed
			// but not yet finalized, without further logic. Once fault proofs are enabled, we will need to update
			// the proposer to wait until the blob associated with an L2 block has been finalized, i.e. the EigenDA
			// contracts on Ethereum have confirmed the full availability of the blob on EigenDA.
			batchHeaderHashHex := fmt.Sprintf("0x%s", hex.EncodeToString(statusRes.Info.BlobVerificationProof.BatchMetadata.BatchHeaderHash))
			fmt.Println("Successfully dispersed blob to EigenDA", "requestID", base64RequestID, "batchHeaderHash", batchHeaderHashHex)
			return statusRes.Info, nil
		} else if statusRes.Status == BlobStatus_UNKNOWN ||
			statusRes.Status == BlobStatus_FAILED {
			fmt.Println("EigenDA blob dispersal failed in processing", "requestID", base64RequestID, "err", err)
			return nil, fmt.Errorf("eigenDA blob dispersal failed in processing with reply status %d", statusRes.Status)
		} else {
			fmt.Println("Still waiting for confirmation from EigenDA", "requestID", base64RequestID)
		}

		// Wait before first status check
		time.Sleep(e.DAStatusQueryRetryInterval)
	}

	return nil, fmt.Errorf("timed out getting EigenDA status for dispersed blob key: %s", base64RequestID)
}
