// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

/**
 * @title MockV3Aggregator
 * @notice 模拟 Chainlink 价格预言机，用于测试
 */
contract MockV3Aggregator is AggregatorV3Interface {
    uint8 private _decimals;
    int256 private _latestAnswer;
    uint256 private _latestTimestamp;
    uint256 private _latestRound;

    string private _description;

    constructor(uint8 decimals_, int256 initialAnswer_) {
        _decimals = decimals_;
        _latestAnswer = initialAnswer_;
        _latestTimestamp = block.timestamp;
        _latestRound = 1;
        _description = "MockV3Aggregator";
    }

    function decimals() external view override returns (uint8) {
        return _decimals;
    }

    function description() external view override returns (string memory) {
        return _description;
    }

    function version() external pure override returns (uint256) {
        return 3;
    }

    function getRoundData(uint80 _roundId) external view override returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    ) {
        return (_roundId, _latestAnswer, _latestTimestamp, _latestTimestamp, _roundId);
    }

    function latestRoundData() external view override returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    ) {
        return (
            uint80(_latestRound),
            _latestAnswer,
            _latestTimestamp,
            _latestTimestamp,
            uint80(_latestRound)
        );
    }

    /**
     * @notice 更新价格（仅测试使用）
     */
    function updateAnswer(int256 answer_) external {
        _latestAnswer = answer_;
        _latestTimestamp = block.timestamp;
        _latestRound++;
    }

    /**
     * @notice 更新价格和时间戳（仅测试使用）
     */
    function updateRoundData(uint80 roundId_, int256 answer_, uint256 timestamp_, uint256 startedAt_) external {
        _latestRound = roundId_;
        _latestAnswer = answer_;
        _latestTimestamp = timestamp_;
    }
}
