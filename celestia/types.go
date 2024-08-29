package celestia

import "github.com/celestiaorg/celestia-node/blob"

type ID struct {
	Height uint64
	SharePointer
	TxHash          string
	ShareCommitment Commitment
}

type Proof = blob.Proof

type Commitment = blob.Commitment

type SharePointer struct {
	Height int64
	Start  int64
	End    int64
}
