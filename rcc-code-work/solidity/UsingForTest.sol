// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

library FindArray {
    function contains(uint[] memory arr, uint i) external pure returns (bool) {
        for (uint x = 0; x < arr.length; x++) {
            if (arr[x] == i) {
                return true;
            }
        }

        return false;
    }
}

contract UsingForTest {
    using FindArray for uint[];

    function find(uint[] memory arr, uint i) public pure returns (bool) {
        return arr.contains(i);
    }
}