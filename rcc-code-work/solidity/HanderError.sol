// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 自定义错误
contract HandlerError {
    error Unauthorized();

    error InsufficionBalance(uint256 available, uint256 _receive);

    error TransferFailed(address from, address to, uint256 amount, string reason);

    function revertError(bool condition) public view returns (uint256) {
        if (!condition) {
            revert TransferFailed(address(this), msg.sender, 10, "20");
        }

        return 1;
    }
}