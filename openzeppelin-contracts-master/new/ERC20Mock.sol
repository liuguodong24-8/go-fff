// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "ERC20.sol";

// mock class using ERC20
contract ERC20Mock is ERC20 {
    constructor(
        string memory name,
        string memory symbol,
        string memory issuerAuthorities,
        string memory issuerName,
        string memory description,
        address initialAccount,
        uint256 initialBalance
    ) payable ERC20(name, symbol, issuerAuthorities, issuerName, description) {
        _mint(initialAccount, initialBalance);
    }

    function mint(address account, uint256 amount) public {
        _mint(account, amount);
    }

    function burn(address account, uint256 amount) public {
        _burn(account, amount);
    }

    function transferInternal(
        address from,
        address to,
        uint256 value
    ) public {
        _transfer(from, to, value);
    }

    function approveInternal(
        address owner,
        address spender,
        uint256 value
    ) public {
        _approve(owner, spender, value);
    }
}