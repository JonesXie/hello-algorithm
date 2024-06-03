package dynamic_programming

import "math"

// 给定n个物品，第i个物品的重量为 wgt[i-1]、价值为 val[i-1]，和一个容量为 cap 的背包。*每个物品可以重复选取*，问在限定背包容量下能放入物品的最大价值。

/* 完全背包：动态规划 */
func unboundedKnapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > c {
				// 若超过背包容量，则不选物品 i
				dp[i][c] = dp[i-1][c]
			} else {
				// 不选和选物品 i 这两种方案的较大值
				dp[i][c] = int(math.Max(float64(dp[i-1][c]), float64(dp[i][c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]
}

/* 完全背包：空间优化后的动态规划 */
func unboundedKnapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([]int, cap+1)
	// 状态转移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > c {
				// 若超过背包容量，则不选物品 i
				dp[c] = dp[c]
			} else {
				// 不选和选物品 i 这两种方案的较大值
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}
