// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

// context
import "../openzeppelin-contracts/contracts/utils/Context.sol";

contract OwnShip is Context{
  address public contract_onwer;

  constructor() {
    contract_onwer = _msgSender();
  }

  modifier OnlyOwner() {
    require(_msgSender() == contract_onwer, "not contract owner");
    _;
  }

  function TransferOwner(address new_onwer) public OnlyOwner{
    contract_onwer = new_onwer;
  }
}