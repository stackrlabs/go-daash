package da

import "context"

// Generic interface to interact with any Data Availability layers.
// Modified fork of https://github.com/rollkit/go-da
type Client interface {
	// MaxBlobSize returns the max blob size
	MaxBlobSize(ctx context.Context) (uint64, error)

	// Get returns Blob for each given ID, or an error.
	//
	// Error should be returned if ID is not formatted properly, there is no Blob for given ID or any other client-level
	// error occurred (dropped connection, timeout, etc).
	Get(ctx context.Context, id ID) (Blob, error)

	// GetProof returns the proof of inclusion for the given ID.
	GetProof(ctx context.Context, id ID) (Proof, error)

	// Commit creates a Commitment for each given Blob.
	Commit(ctx context.Context, blob Blob) (Commitment, error)

	// Submit submits the Blobs to Data Availability layer.
	//
	// This method is synchronous. Upon successful submission to Data Availability layer, it returns ID identifying blob
	// in DA and Proof of inclusion.
	// If options is nil, default options are used.
	Submit(ctx context.Context, blob Blob, gasPrice float64) (ID, error)

	// Validate validates Commitments against the corresponding Proofs. This should be possible without retrieving the Blobs.
	Validate(ctx context.Context, id ID, proof Proof) (bool, error)
}

// Blob is the data submitted/received from DA interface.
type Blob []byte

// ID should hold data required by the implementation to find blob in Data Availability layer.
type ID any

// Commitment should contain cryptographic commitment to Blob value.
type Commitment any

// Proof should contain a proof of inclusion (publication) of Blob in Data Availability layer.
type Proof any
