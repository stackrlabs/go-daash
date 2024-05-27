// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import {DAVerifier, SharesProof, NamespaceNode, BinaryMerkleProof, AttestationProof} from "blobstream-contracts/lib/verifier/DAVerifier.sol";
import {IDAOracle} from "blobstream-contracts/IDAOracle.sol";

/// @notice A span sequence defines the location of the rollup transaction data
/// inside the Celestia block.
/// For this tutorial, we will be posting the rollup data to a single block. Thus, all
/// we will need is the height, the index and the length of the data.
/// This can be generalized to multiple Celestia blocks.
struct SpanSequence {
    // Celestia block height where the rollup data was posted.
    uint256 height;
    // Index of the first share containing the rollup transaction data
    // inside the Celestia block
    uint256 index;
    // Number of shares that the rollup transaction data spans on.
    uint256 length;
}

contract BlobstreamVerifier {
    // Verify data availability
    function verifyDataAvailability(
        address bridge,
        SpanSequence memory span,
        NamespaceNode[] memory _rowRoots,
        BinaryMerkleProof[] memory _rowProofs,
        AttestationProof memory _attestationProof,
        bytes32 root
    ) public view returns (bool) {
        // 1. Verify availability of some shares (corresponding to row root) to a Celestia-attested block range
        (bool res, ) = DAVerifier.verifyMultiRowRootsToDataRootTupleRoot(
            IDAOracle(bridge),
            _rowRoots,
            _rowProofs,
            _attestationProof,
            root
        );
        require(res, "Row root to data root tuple root verification failed");

        // 2. Verify if shares in input span belongs to the above authenticated row root
        // 2a. Make sure span matches block height in attestation proof
        require(span.height == _attestationProof.tuple.height, "Span does not correspond to block height of the row root");

        // 2b. Make sure span startIndex is in the share range of the first row root we just authenticated
        (uint256 squareSize, DAVerifier.ErrorCodes error) =
            DAVerifier.computeSquareSizeFromRowProof(_rowProofs[0]);
        require(uint8(error) == uint8(DAVerifier.ErrorCodes.NoError), "Error computing square size");
        uint256 rowStartIndex = squareSize * _rowProofs[0].key;
        require(span.index >= rowStartIndex, "Data is unavailable");

        // 2c. Make sure span startIndex is in the share range of the last row root we just authenticated
        uint256 rowEndIndex = squareSize * (_rowProofs[_rowProofs.length - 1].key + 1);
        uint256 endIndex = span.index + span.length;
        require(endIndex <= rowEndIndex, "Data is unavailable");

        return true;  
    }

    function verifyDataInclusion(
        address bridge,
        SharesProof memory sharesProof, // raw share data contained in this
        bytes32 root,
        SpanSequence memory span // Pointer to a blob on DA. We assume its availability has already been proven.
    ) public view returns (bool) {
        // Verify share is part of an attested block range
        (bool success, ) = DAVerifier.verifySharesToDataRootTupleRoot(
            IDAOracle(bridge),
            sharesProof,
            root
        );
        require(success, "Shares to data root tuple root verification failed");

        // This steo proves that the share is part of the provided SpanSequence
        // To do so, we will use the proof, already authenticated above, to get the index,
        // Then, we will compare it against the spans sequence.
        // since we're using nmt multiproofs, we have a begin key and an end key of the shares
        // proven. However, in our case, we're only proving a single share.
        // Thus, we can take the begin key as the index.
        // Note: In the case of multiple shares in the proof, we will need to check all the shares
        // if they're part of the sequence of spans. Then, only use the ones that are part of it.
        uint256 endIndex = span.index + span.length;
        uint256 shareIndexInRow = sharesProof.shareProofs[0].beginKey;
        uint256 shareIndexInRowMajorOrder =
            shareIndexInRow + (sharesProof.rowProofs[0].numLeaves / 4) * sharesProof.rowProofs[0].key;

        // check if the share is part of the sequence of spans
        require(span.index <= shareIndexInRowMajorOrder, "Share is not part of the sequence of spans");
        require(shareIndexInRowMajorOrder <= endIndex, "Share is not part of the sequence of spans");

        return true;
    }
}
