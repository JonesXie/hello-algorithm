package dynamic_programming

import "math"

// 背包问题是一大类动态规划问题的代表，其拥有很多变种，例如零钱兑换问题。

// 给定n种物品，第i种硬币的面值为 coins[i-1]、目标金额为 amt。*每个硬币可以重复选取*，问能够凑出目标金额的最少硬币数量。如果无法凑出目标金额，则返回 -1

/* 零钱兑换：动态规划 */
func coinChangeDP(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}
	// 状态转移：首行首列
	for a := 1; a <= amt; a++ {
		dp[0][a] = max
	}
	// 状态转移：其余行和列
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				dp[i][a] = dp[i-1][a]
			} else {
				// 不选和选硬币 i 这两种方案的较小值
				dp[i][a] = int(math.Min(float64(dp[i-1][a]), float64(dp[i][a-coins[i-1]]+1)))
			}
		}
	}
	if dp[n][amt] != max {
		return dp[n][amt]
	}
	return -1
}

/* 零钱兑换：动态规划 */
func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	// 初始化 dp 表
	dp := make([]int, amt+1)
	for i := 1; i <= amt; i++ {
		dp[i] = max
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		// 正序遍历
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				dp[a] = dp[a]
			} else {
				// 不选和选硬币 i 这两种方案的较小值
				dp[a] = int(math.Min(float64(dp[a]), float64(dp[a-coins[i-1]]+1)))
			}
		}
	}
	if dp[amt] != max {
		return dp[amt]
	}
	return -1
}
