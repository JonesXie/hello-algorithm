package greedy

import "math"

// 给定一个正整数 n ，将其切分为至少两个正整数的和，求切分后所有整数的乘积最大是多少

/* 最大切分乘积：贪心 */
func maxProductCutting(n int) int {
	// 当 n <= 3 时，必须切分出一个 1
	if n <= 3 {
		return 1 * (n - 1)
	}
	// 贪心地切分出 3 ，a 为 3 的个数，b 为余数
	a := n / 3
	b := n % 3
	if b == 1 {
		// 当余数为 1 时，将一对 1 * 3 转化为 2 * 2
		return int(math.Pow(3, float64(a-1))) * 2 * 2
	}
	if b == 2 {
		// 当余数为 2 时，不做处理
		return int(math.Pow(3, float64(a))) * 2
	}
	// 当余数为 0 时，不做处理
	return int(math.Pow(3, float64(a)))
}
