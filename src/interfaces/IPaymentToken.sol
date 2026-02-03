// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

/**
 * @title IPaymentToken
 * @notice 支付代币接口，用于统一 ETH 和 ERC20 代币的支付处理
 * @dev 为未来扩展到多种 ERC20 代币（USDC, DAI 等）提供基础
 */
interface IPaymentToken {
    /**
     * @notice 支付方式枚举
     * @dev ETH: 原生以太坊，WETH: Wrapped Ethereum
     * 预留位置供未来添加 USDC, DAI 等
     */
    enum PaymentMethod {
        ETH,    // 原生 ETH
        WETH    // Wrapped ETH
        // 未来扩展: USDC, DAI, USDT 等
    }

    /**
     * @notice 获取支付方式对应的代币地址
     * @param method 支付方式
     * @return 代币地址，ETH 返回 address(0)
     */
    function getPaymentTokenAddress(PaymentMethod method) external view returns (address);

    /**
     * @notice 检查支付方式是否支持
     * @param method 支付方式
     * @return 是否支持
     */
    function isPaymentMethodSupported(PaymentMethod method) external view returns (bool);
}
