// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ERC721Demo1 is ERC721, ERC721URIStorage, ERC721Enumerable, Ownable {
    // 如果需要最大发行量就加上
    uint256 public immutable maxSupply;

    constructor(uint256 _maxSupply) 
        ERC721("RTT", "RTT") 
        Ownable(msg.sender) 
    {
        require(_maxSupply > 0, "Invalid max supply");
        maxSupply = _maxSupply;
    }

    // 核心铸造功能
    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 currentSupply = totalSupply();
        require(currentSupply < maxSupply, "Max supply exceeded");
        
        _safeMint(to, currentSupply); // 使用当前总量作为新 tokenId
        _setTokenURI(currentSupply, uri);
    }

    // 新增的枚举功能接口
    function totalSupply() public view override(ERC721Enumerable) returns (uint256) {
        return super.totalSupply();
    }

    function tokenByIndex(uint256 index) public view override returns (uint256) {
        return super.tokenByIndex(index);
    }

    function tokenOfOwnerByIndex(address owner, uint256 index) public view override returns (uint256) {
        return super.tokenOfOwnerByIndex(owner, index);
    }

    // 多重继承需要的方法覆盖
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721, ERC721Enumerable)
    {
        super._increaseBalance(account, value);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage, ERC721Enumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}