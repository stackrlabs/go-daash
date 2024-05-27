// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {VectorVerifier} from "../src/VectorVerifier.sol";

contract DeployScript is Script {
    function setUp() public {}
    
    function run() public returns (address) {
        vm.startBroadcast();
        VectorVerifier verifier = new VectorVerifier();
        vm.stopBroadcast();
        return address(verifier);
    }
}
