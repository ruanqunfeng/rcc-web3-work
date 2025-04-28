// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract Voting {
    mapping (string name => uint num) votes;

    function vote(string calldata username)  public{
        votes[username]++;
    }

    function getVotes(string calldata username) public view returns(uint){
        return votes[username];
    }

    function resetVotes(string calldata username) public {
        votes[username] = 0;
    }
}