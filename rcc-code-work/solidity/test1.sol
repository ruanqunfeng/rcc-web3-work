// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract Hello3Dot1 {

    // 基本类型，存储在Stack中
    int public x = 1;

    uint public ux = 0;

    bool public b = true;

    address public addr = 0x4aFc74781822F4B1eCa779312458b7EE75D7B69a;

    bytes2 public b2 = 0x1210;

    bytes32 public b32 = hex"10";

    enum Color { Red, Green, Blue }

    function getX() public view returns (int) {
        return x;
    }

    // 引用类型，存储在Memory中
    int[] public  arr = [int(1), 2, 3];

    string public name = "Hello3Dot1";

    struct Person {
        int age;
        string name;
        //mapping (address => uint) m1;
    }

    Person public person = Person(10, "Alice");
    Person public person2 = Person({age: 20, name: "Bob"});

    Person[] public persons;
    function addPerson(int age, string memory name1) public {
        persons.push(Person(age, name1));
    }

    bytes public bs = "Hello3Dot1";
    function getBs(uint256 i) public view returns (bytes1) {
        return bs[i];
    }

    function testSlice(bytes calldata bs1) public pure returns (bytes memory) {
        return bs1[0:2];
    }
}