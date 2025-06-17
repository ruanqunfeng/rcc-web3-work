// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";

contract OZInterfaceChecker {
    using ERC165Checker for address;

    /// ERC‑721 与 ERC‑1155 的 interface IDs
    bytes4 public constant ERC721_ID  = 0x80ac58cd;
    bytes4 public constant ERC1155_ID = 0xd9b67a26;

    /// 检测
    function isERC721(address target) external view returns (bool) {
        return target.supportsInterface(ERC721_ID);
    }

    function isERC1155(address target) external view returns (bool) {
        return target.supportsInterface(ERC1155_ID);
    }

    /// 通用检测
    function isSupports(address target, bytes4 interfaceId) external view returns (bool) {
        return target.supportsInterface(interfaceId);
    }
}
