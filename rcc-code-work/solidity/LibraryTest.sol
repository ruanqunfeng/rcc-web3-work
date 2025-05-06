// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./AddTest.sol";
contract LibraryTest {
    function add(uint x, uint y) public pure returns (uint) {
        return SafeMath.add(x, y);  // 内嵌库函数调用
    }
}