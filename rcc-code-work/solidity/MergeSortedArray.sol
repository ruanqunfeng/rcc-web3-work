// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract MergeSortedArray {
    function merge(int256[] memory arr1, int256[] memory arr2)
        public
        pure
        returns (int256[] memory)
    {
        int256[] memory arr3 = new int256[](arr1.length + arr2.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 index = 0;
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] <= arr2[j]) {
                arr3[index] = arr1[i];
                index++;
                i++;
            } else {
                arr3[index] = arr2[j];
                index++;
                j++;
            }
        }

        while (i < arr1.length) {
            arr3[index] = arr1[i];
            index++;
            i++;
        }

        while (j < arr2.length) {
            arr3[index] = arr2[j];
            index++;
            j++;
        }

        return arr3;
    }
}