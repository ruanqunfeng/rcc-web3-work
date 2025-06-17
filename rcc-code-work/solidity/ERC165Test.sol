// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/interfaces/IERC165.sol";

contract ERC165Test is IERC165 {

    // 直接写入函数的选择器
    bytes4 private constant _INTERFACE_ID_ERC165 = 0x01ffc9a7;

    bytes4 private constant _TEST_ID = bytes4(keccak256("test()"));

    mapping(bytes4 => bool) private supported;

    constructor() {
        _registerInterface(_INTERFACE_ID_ERC165);
        _registerInterface(_TEST_ID);
    }

    function supportsInterface(bytes4 interfaceId) public view returns (bool) {
        return supported[interfaceId];
    }

    function _registerInterface(bytes4 interfaceId) internal {
        require(interfaceId != 0xffffffff, "ERC165: invalid interface id");
        supported[interfaceId] = true;
    }

    function test() public pure {

    }

}