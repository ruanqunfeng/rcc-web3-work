// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract ReverseString {
    function reverse(string memory str) public pure returns (string memory) {
        bytes memory bs = bytes(str);
        bytes memory bs2 = new bytes(bs.length);
        for (uint i = 0; i < bs.length; i++) {
            bs2[i] = bs[bs.length - 1 - i];
        }

        return string(bs2);
    }
}