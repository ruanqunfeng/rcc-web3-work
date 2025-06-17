// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract InterfaceChecker {
    bytes4 private constant _INTERFACE_ID_ERC165 = 0x01ffc9a7;
    bytes4 private constant _INTERFACE_ID_ERC721 = 0x80ac58cd;
    bytes4 private constant _INTERFACE_ID_ERC1155 = 0xd9b67a26;
    bytes4 private constant ERC165_SELECTOR = 0x01ffc9a7; // supportsInterface 的函数选择器

    function isContract(address addr) private view returns(bool) {
        uint256 size;
        assembly { size := extcodesize(addr) }
        return size > 0;
    }

    // 检查是否为 ERC721 合约
    function isERC721(address _contract) public view returns (bool) {
        return supportsInterface(_contract, _INTERFACE_ID_ERC721);
    }

    // 检查是否为 ERC1155 合约
    function isERC1155(address _contract) public view returns (bool) {
        return supportsInterface(_contract, _INTERFACE_ID_ERC1155);
    }

    // 通用接口检查函数
    function supportsInterface(address _contract, bytes4 interfaceId) public view returns (bool) {
        // 先检查是否是合约地址
        if (!isContract(_contract)) {
            return false;
        }

        // 先检查是否支持 ERC165
        if (!_supportsERC165(_contract)) {
            return false;
        }

        // 检查目标接口支持情况
        (bool success, bytes memory result) = _contract.staticcall{gas: 30000}(
            abi.encodeWithSelector(ERC165_SELECTOR, interfaceId)
        );
        
        return success && result.length >= 32 && abi.decode(result, (bool));
    }

    // 内部 ERC165 检查
    function _supportsERC165(address _contract) internal view returns (bool) {
        (bool success, bytes memory result) = _contract.staticcall{gas: 30000}(
            abi.encodeWithSelector(ERC165_SELECTOR, _INTERFACE_ID_ERC165)
        );
        
        return success && result.length >= 32 && abi.decode(result, (bool));
    }
}