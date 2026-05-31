// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockUSDC is ERC20 {
    // 真实 USDC 为 6 位精度
    uint8 private constant _DECIMALS = 6;

    constructor() ERC20("MockUSDC", "USDC") {
        // 初始给部署者铸造 1,000,000 USDC（6 位精度）
        _mint(msg.sender, 1_000_000 * 10 ** _DECIMALS);
    }

    function decimals() public view virtual override returns (uint8) {
        return _DECIMALS;
    }

    // 任何人都可以调，方便测试
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}