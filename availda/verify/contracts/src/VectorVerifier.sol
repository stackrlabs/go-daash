// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import {IAvailBridge} from "./interfaces/IAvailBridge.sol";

contract VectorVerifier {
    function verifyDataAvailability(
        address bridge,
        bytes32 blobHash, 
        IAvailBridge.MerkleProofInput memory input
    ) public view returns (bool) {
        // Verify blob Hash correspnds to leaf in merkle proof
        require(input.leaf == blobHash, "Blob hash does not correspond to leaf in merkle proof");
        //  Verify availability of blob corresponding to leaf in merkle proof
        IAvailBridge availBridge = IAvailBridge(bridge);
        bool success = availBridge.verifyBlobLeaf(input);
        return success;
    }

    function verifyDataInclusion(
        address bridge,
        bytes memory blobData,
        IAvailBridge.MerkleProofInput memory input
    ) public view returns (bool) {
        // Hash the blob data to get the blob hash
        bytes32 blobHash = keccak256(blobData);
        // Verify the blob is available
        return verifyDataAvailability(bridge, blobHash, input);
    }
}
