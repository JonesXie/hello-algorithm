package dynamic_programming

// 给定一个楼梯，你每步可以上 1 阶或者 2 阶，每一阶楼梯上都贴有一个非负整数，表示你在该台阶所需要付出的代价。
//给定一个非负整数数组  cost，其中 cost[i] 表示在第 i 个台阶需要付出的代价， cost[0] 为地面（起始点）。请计算最少需要付出多少代价才能到达顶部？

/* 爬楼梯最小代价：动态规划 */
func minCostClimbingStairsDP(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([]int, n+1)
	// 初始状态：预设最小子问题的解
	dp[1] = cost[1]
	dp[2] = cost[2]
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]
}

/* 爬楼梯最小代价：空间优化后的动态规划 */
func minCostClimbingStairsDPComp(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 初始状态：预设最小子问题的解
	a, b := cost[1], cost[2]
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		tmp := b
		b = min(a, tmp) + cost[i]
		a = tmp
	}
	return b
}
