// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


import './token/ERC20/IERC20.sol';

import "./utils/Counters.sol";

/// @notice This is a token manage contract
/// @dev test
contract Manage{
    address private token;
    address private admin;
    
    mapping(address=>uint256) private balance;

    modifier onlyAdmin(){
         require(admin == msg.sender, "Manage: E01");
        _;
    }

    constructor(address _admin, address _token) {
        admin = _admin;
        token = _token;
	}

    function transferToken(address from, address to, uint256 value) external onlyAdmin{
        require(balance[from] >= value, "Manage:E01");
        balance[to] += value;
        balance[from] -= value;
    }

    function deposit(address from, uint256 value) external onlyAdmin{
        require( IERC20(token).balanceOf(from) >= value, "Manage:E02" );
        balance[from] += value;
        IERC20(token).transferFrom(from, address(this), value);
    }

    function withdraw(address from, uint256 value) external onlyAdmin{
        require( balance[from] >= value, "Manage:E03" );
        require( IERC20(token).balanceOf(address(this)) >= value, "Manage:E03" );

        balance[from] -= value;
        IERC20(token).transfer(from, value);
    }

    function balanceToken(address from) external view returns(uint256 value){
        value = balance[from];
    }
}