// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SimpleCAW is Ownable {
    IERC20 public usdc;

    struct RecipientPolicy {
        uint256 dailyLimit;      // 每天最多给这个地址付多少
        uint256 spentToday;      // 今天已付多少
        uint256 lastResetDay;    // 上次重置是哪一天（block.timestamp / 1 days）
    }

    mapping(address => RecipientPolicy) public policies;

    event PaymentExecuted(address indexed recipient, uint256 amount);
    event PolicyUpdated(address indexed recipient, uint256 dailyLimit);

    constructor(address usdc_) Ownable(msg.sender) {
        usdc = IERC20(usdc_);
    }

    // 设置某个收款方的策略（模拟 Pact / Policy）
    function setLimit(address recipient, uint256 dailyLimit) external onlyOwner {
        policies[recipient] = RecipientPolicy({
            dailyLimit: dailyLimit,
            spentToday: 0,
            lastResetDay: block.timestamp / 1 days
        });
        emit PolicyUpdated(recipient, dailyLimit);
    }

    // 模拟 Agent 通过 CAW 执行支付：在策略允许范围内转账
    function executePayment(address recipient, uint256 amount) external onlyOwner {
        RecipientPolicy storage policy = policies[recipient];

        // 1. 白名单检查：如果 dailyLimit == 0，视为未配置/不允许
        require(policy.dailyLimit > 0, "CAW: recipient not in whitelist");

        // 2. 按天重置计数
        uint256 today = block.timestamp / 1 days;
        if (policy.lastResetDay < today) {
            policy.spentToday = 0;
            policy.lastResetDay = today;
        }

        // 3. 预算检查
        require(
            policy.spentToday + amount <= policy.dailyLimit,
            "CAW: exceeds daily budget limit"
        );

        // 4. 执行 USDC 转账
        bool ok = usdc.transfer(recipient, amount);
        require(ok, "CAW: USDC transfer failed");

        // 5. 更新已用额度
        policy.spentToday += amount;

        emit PaymentExecuted(recipient, amount);
    }
}
