// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract MyContract {
    struct MyStruct {
        uint256 id;
        mapping(string => uint256) myMapping;
    }

    MyStruct myStruct;

    function setMapping(string memory key, uint256 value) public {
        myStruct.myMapping[key] = value;
    }

    function getMapping(string memory key) public view returns (uint256) {
        return myStruct.myMapping[key];
    }
}