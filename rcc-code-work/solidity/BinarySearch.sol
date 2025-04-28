// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {

    function binarySearch(int[] memory arr, int target) public pure returns (int) {
        uint i = 0;
        // j有可能是负数，所以要转成int
        int j = int(arr.length) - 1;
        while (int(i) <= j) {
            uint mid = i + ((uint(j) - i) / 2);
            if (arr[mid] == target) {
                return int(mid);
            } else if (arr[mid] < target ) {
                i = mid + 1;
            } else {
                j = int(mid) - 1;
            }
        }

        return -1;
    }
}