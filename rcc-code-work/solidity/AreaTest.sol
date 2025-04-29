// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract AreaTest {
    uint public i1 = 10;
    uint internal i2 = 20;
    uint private i3 = 30;

    function foo1() public view returns (uint) {
        return i3;
    }
}

contract Test1 {
    AreaTest area;
    
    constructor(address _area_addr){
        area = AreaTest(_area_addr); // 构造函数中调用构造函数，并且传递参数给构造函数。在 contract 内部调用的是外部合约的构造函数，这样就能够在 contract 内部直接调用外部合约的函数
        // 异常，不能直接调用父类合约的变量
        // uint _i = area.i2();
    }

}

contract Test2 is AreaTest {
    AreaTest area;
    
    // 异常，不能声明父类合约的变量
    // uint public i2 = 20;

    constructor(address _area_addr) {
        area = AreaTest(_area_addr); // 构造函数中调用构造函数，并且传递参数给构造函数。在 contract 内部调用的是外部合约的构造函数，这样就能够在 contract 内部直接调用外部合约的函数
        // 可以直接调用父类的变量
        uint _i = i2;
    }
}


contract ImmutableExample {
    uint immutable maxBalance;

    constructor(uint _maxBalance) {
        maxBalance = _maxBalance;
    }

    uint constant x = 42;

    function test1(uint _x) public {
        // x = _x;
    }
}