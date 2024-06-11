package verify

import (
	"context"
	"fmt"

	"github.com/stackrlabs/go-daash/celestia"
	bv "github.com/stackrlabs/go-daash/celestia/verify/blobstream/verifier"
)

func (v Verifier) GetShareProof(sp celestia.SharePointer) (*bv.SharesProof, [32]byte, error) {
	ctx := context.Background()

	// 1. Get the data commitment
	dataCommitment, err := GetDataCommitment(v.ethClient, sp.Height, 10_000_000, v.blobstreamXContract)
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to get data commitment: %w", err)
	}

	// 2. Get the block
	blockRes, err := v.tRPCClient.Block(ctx, &sp.Height)
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to get block: %w", err)
	}

	// 3. get data root inclusion commitment
	dcProof, err := v.tRPCClient.DataRootInclusionProof(ctx, uint64(sp.Height), dataCommitment.StartBlock, dataCommitment.EndBlock)
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to get data root inclusion proof: %w", err)
	}

	// 4. get share proof
	shareProof, err := v.tRPCClient.ProveShares(ctx, (uint64(sp.Height)), uint64(sp.Start), uint64(sp.End))
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to get share proof: %w", err)
	}

	nonce := dataCommitment.ProofNonce.Uint64()
	height := uint64(sp.Height)

	blockDataRoot := [32]byte(blockRes.Block.DataHash)

	return &bv.SharesProof{
		Data:             shareProof.Data,
		ShareProofs:      toNamespaceMerkleMultiProofs(shareProof.ShareProofs),
		Namespace:        *namespace(shareProof.NamespaceID, uint8(shareProof.NamespaceVersion)),
		RowRoots:         toRowRoots(shareProof.RowProof.RowRoots),
		RowProofs:        toRowProofs(shareProof.RowProof.Proofs),
		AttestationProof: toAttestationProof(nonce, height, blockDataRoot, dcProof.Proof),
	}, blockDataRoot, nil
}
