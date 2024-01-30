package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type EigendaDAClient struct {
	// DaRpc is the HTTP provider URL for the Data Availability node.
	DARpc string

	// Quorum IDs and SecurityParams to use when dispersing and retrieving blobs
	DADisperserSecurityParams []*SecurityParams

	// The total amount of time that the batcher will spend waiting for EigenDA to confirm a blob
	DAStatusQueryTimeout time.Duration

	// The amount of time to wait between status queries of a newly dispersed blob
	DAStatusQueryRetryInterval time.Duration
}

func NewEigendaDAClient(daRpc string, daStatusQueryTimeout time.Duration, daStatusQueryRetryInterval time.Duration) *EigendaDAClient {
	disperserSecurityParams := []*SecurityParams{}
	disperserSecurityParams = append(disperserSecurityParams, &SecurityParams{
		QuorumId:           0,
		AdversaryThreshold: 25,
		QuorumThreshold:    50,
	})
	return &EigendaDAClient{
		DARpc:                      daRpc,
		DADisperserSecurityParams:  disperserSecurityParams,
		DAStatusQueryTimeout:       daStatusQueryTimeout,
		DAStatusQueryRetryInterval: daStatusQueryRetryInterval,
	}
}

func (e *EigendaDAClient) disperseBlob(ctx context.Context, txData []byte) (*BlobInfo, error) {
	fmt.Println("Attempting to disperse blob to EigenDA")
	// Use secure credentials
	conn, err := grpc.Dial(e.DARpc, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		fmt.Printf("Unable to connect to EigenDA, aborting", "err", err)
		return nil, err
	}
	defer conn.Close()
	daClient := NewDisperserClient(conn)

	disperseReq := &DisperseBlobRequest{
		Data:           txData,
		SecurityParams: e.DADisperserSecurityParams,
	}
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
