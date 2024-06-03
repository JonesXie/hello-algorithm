package dynamic_programming

// 暴力搜索

/* 搜索 */
func dfs(i int) int {
	// 已知 dp[1] 和 dp[2] ，返回之
	if i == 1 || i == 2 {
		return i
	}
	// dp[i] = dp[i-1] + dp[i-2]
	count := dfs(i-1) + dfs(i-2)
	return count
}

/* 爬楼梯：搜索 */
func climbingStairsDFS(n int) int {
	return dfs(n)
}
