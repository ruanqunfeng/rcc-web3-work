// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract SimpleWallet  {
    address payable receiveAddr = payable(0x4aFc74781822F4B1eCa779312458b7EE75D7B69a);
    uint balance;

    // 往合约转账
    function deposit() public payable {
        balance += msg.value;
    }

    function getBalance() public view returns(uint) {
        return balance;
    }

    // 提现
    function withdraw(uint _amount) public returns(bool) {
        require(balance >= _amount, unicode"余额不足");
        balance -= _amount;
        bool success = receiveAddr.send(_amount);
        if (!success) {
            balance += _amount;
        }

        return success;
    }

    uint[] public arr = [1,2,3];

    function getReceiveAddrBalance() public view returns(uint) {
        return receiveAddr.balance;
    }

    uint[] dataMemory = new uint[](7);
}