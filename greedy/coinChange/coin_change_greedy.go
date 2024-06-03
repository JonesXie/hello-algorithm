package greedy

// 给定n种硬币，第i种硬币的面值为 coins[i-1]、目标金额为 amt。*每个硬币可以重复选取*，问能够凑出目标金额的最少硬币数量。如果无法凑出目标金额，则返回 -1

/* 零钱兑换：贪心 */
func coinChangeGreedy(coins []int, amt int) int {
	// 假设 coins 列表有序
	i := len(coins) - 1
	count := 0
	// 循环进行贪心选择，直到无剩余金额
	for amt > 0 {
		// 找到小于且最接近剩余金额的硬币
		for i > 0 && coins[i] > amt {
			i--
		}
		// 选择 coins[i]
		amt -= coins[i]
		count++
	}
	// 若未找到可行方案，则返回 -1
	if amt != 0 {
		return -1
	}
	return count
}
