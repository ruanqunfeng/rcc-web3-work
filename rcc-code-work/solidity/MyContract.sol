// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract MyContract {

    bytes10 bs = unicode"中国龙1";

    bytes public array10 = hex"11";

    function getArray10Len() public view returns (uint){
        return array10.length;
    }

    struct MyStruct {
        uint256 id;
        mapping(string => uint256) myMapping;
    }

     function test() internal returns(MyStruct storage) {
        return myStruct;
    }

    struct A {
        uint256 id;
    }

    MyStruct myStruct;

    function setMapping(string calldata key, uint256 value) public {
        myStruct.myMapping[key] = value;
    }

    function getMapping(string memory key) public view returns (uint256) {
        return myStruct.myMapping[key];
    }

    // 0000 0010
    // 1111 1110
    function sub1() public pure returns (uint8) { 
        uint8 m = 1;
        uint8 n = m - 2; 
        return n; 
    } 
}