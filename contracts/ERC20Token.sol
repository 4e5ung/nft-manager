// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./token/ERC20/ERC20.sol";
import "./token/ERC20/utils/SafeERC20.sol";
import "./access/Ownable.sol";

contract ERC20Token is ERC20, Ownable{
    using SafeERC20 for IERC20;

    event Burned(address request, uint256 amount);
    
    constructor() ERC20("ERC20 Token", "TEST") {
        _mint(msg.sender, (10 ** 9) * (10 ** 18));
    }

    function burn(uint256 _amount) external onlyOwner returns (bool) {
        _burn(msg.sender, _amount);
        emit Burned(msg.sender, _amount);
        return true;
    }
}