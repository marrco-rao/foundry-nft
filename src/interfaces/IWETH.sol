// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IWETH
 * @notice Wrapped Ethereum (WETH) 接口
 * @dev 扩展自 IERC20，增加 ETH 包装和解包装功能
 */
interface IWETH is IERC20 {
    /**
     * @notice 存入 ETH 并获得等量 WETH
     */
    function deposit() external payable;

    /**
     * @notice 销毁 WETH 并取回等量 ETH
     * @param amount 要解包装的 WETH 数量
     */
    function withdraw(uint256 amount) external;
}
