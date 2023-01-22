// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./lib.sol";

contract Test1 {
    function test() external pure returns (uint) {
        return Lib.f();
    }
}
