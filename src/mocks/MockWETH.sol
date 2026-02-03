// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../interfaces/IWETH.sol";

/**
 * @title MockWETH
 * @notice Wrapped Ethereum (WETH) 模拟合约，用于测试
 * @dev 实现 IWETH 接口，支持 ETH 包装和解包装
 */
contract MockWETH is ERC20, IWETH {
    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    constructor() ERC20("Wrapped Ether", "WETH") {}

    /**
     * @notice 存入 ETH 并获得等量 WETH
     */
    function deposit() public payable override {
        _mint(msg.sender, msg.value);
        emit Deposit(msg.sender, msg.value);
    }

    /**
     * @notice 销毁 WETH 并取回等量 ETH
     * @param amount 要解包装的 WETH 数量
     */
    function withdraw(uint256 amount) public override {
        require(balanceOf(msg.sender) >= amount, "Insufficient WETH balance");
        _burn(msg.sender, amount);
        payable(msg.sender).transfer(amount);
        emit Withdrawal(msg.sender, amount);
    }

    /**
     * @notice 允许合约接收 ETH
     */
    receive() external payable {
        deposit();
    }
}
