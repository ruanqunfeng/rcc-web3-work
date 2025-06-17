// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20 {

    constructor() ERC20("CTT", "CTT") {
        _mint(msg.sender, 1000 * 10 ** 18);  // 初始化账户，给合约部署者 1000 代币
    }

    function checkBalance(address account) public view returns (uint256) {
        return balanceOf(account);  // 返回指定账户的代币余额
    }

    function sendTokens(address recipient, uint256 amount) public returns (bool) {
        return transfer(recipient, amount);  // 将代币发送给接收者
    }

    // 授权第三方可以使用代币
    function approveTokens(address spender, uint256 amount) public returns (bool) {
        return approve(spender, amount);  // 授权 spender 可以使用 amount 个代币
    }

    // 查询剩余的授权额度
    function checkAllowance(address owner, address spender) public view returns (uint256) {
        return allowance(owner, spender);  // 查询 owner 授权给 spender 的额度
    }

    // 第三方（被授权者）代替 owner 使用代币
    function transferFromOwner(address owner, address recipient, uint256 amount) public returns (bool) {
        return transferFrom(owner, recipient, amount);  // 代表 owner 将代币转移给 recipient
    }
}