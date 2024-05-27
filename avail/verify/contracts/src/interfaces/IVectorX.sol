// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

interface IVectorx {
    function dataRootCommitments(bytes32 rangeHash) external view returns (bytes32 dataRoot);
    function rangeStartBlocks(bytes32 rangeHash) external view returns (uint32 startBlock);
}