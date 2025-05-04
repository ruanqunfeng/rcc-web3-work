// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ABI {
    function encodeData(string memory text, uint256 number) public pure returns (bytes memory, bytes memory) {
        return (abi.encode(text, number), abi.encodePacked(text, number));
    }

    function decodeData(bytes memory bs) public pure returns (string memory, uint256) {
        return (abi.decode(bs, (string, uint256)));
    }

    // 函数选择器
    function getSelector() public pure returns (bytes4) {
        return msg.sig;
    }

    // 通过函数名称获取函数选择器
    function computeSelector(string memory str) public pure returns (bytes4) {
        return bytes4(keccak256(bytes(str)));
    }

    function transfer(address addr, uint256 amount) public pure returns(bytes memory) {
        return msg.data;
    }

    // 相当于调用了transfer方法，一定要去掉空格
    function encodeFunctionCall() public pure returns(bytes memory) {
        return abi.encodeWithSignature("transfer(address,uint256)", 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2, 10);
    }

    // hash运算
    function hashFunctions(string memory str) public pure returns(bytes32, bytes32, bytes32) {
        bytes memory bs = abi.encodePacked(str);
        return (keccak256(bs), sha256(bs), ripemd160(bs));
    }

    // 数学运算
    function mathTest(uint256 x, uint256 y, uint256 m) public pure returns (uint256, uint256) {
        // (x + y) % m
        // (x * y) % m
        return (addmod(x,y,m), mulmod(x, y, m));
    }

    // 椭圆曲线恢复公钥
    function recoverAddress(bytes32 hash, uint8 v, bytes32 r, bytes32 s) public pure returns(address) {
        return ecrecover(hash, v, r, s);
    }

    // 调用地址的代码
    function callContract(address addr, bytes memory data) public returns (bool, bytes memory) {
        (bool success, bytes memory result) = addr.call(data);
        return (success, result);
    }
}