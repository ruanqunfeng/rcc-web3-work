// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

abstract contract Contract1 {
    function test() public virtual payable;
}


abstract contract Animal {
    // 抽象函数（无实现）
    function speak() public virtual returns (string memory);
    
    // 已实现的函数
    function eat() public pure returns (string memory) {
        return "Eating";
    }
}

contract Dog is Animal, Contract1 {

    constructor() { 

    }

    function speak() public pure override returns (string memory) {
        return "Bark";
    }

    function test() public override payable { }
}