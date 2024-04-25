package main

import "fmt"

// 分治搜索策略

// 给定一个长度为 n 的有序数组 nums ，其中所有元素都是唯一的，请查找元素 target 。

/* 二分查找：问题 f(i, j) */
func dfs(nums []int, target, i, j int) int {
	// 如果区间为空，代表没有目标元素，则返回 -1
	if i > j {
		return -1
	}
	//    计算索引中点
	m := i + ((j - i) >> 1)
	//判断中点与目标元素大小
	if nums[m] < target {
		// 小于则递归右半数组
		// 递归子问题 f(m+1, j)
		return dfs(nums, target, m+1, j)
	} else if nums[m] > target {
		// 小于则递归左半数组
		// 递归子问题 f(i, m-1)
		return dfs(nums, target, i, m-1)
	} else {
		// 找到目标元素，返回其索引
		return m
	}
}

/* 二分查找 */
func binarySearch(nums []int, target int) int {
	n := len(nums)
	return dfs(nums, target, 0, n-1)
}

func main() {
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	target := 6
	noTarget := 99
	targetIndex := binarySearch(nums, target)
	fmt.Println("目标元素 6 的索引 = ", targetIndex)
	noTargetIndex := binarySearch(nums, noTarget)
	fmt.Println("不存在目标元素的索引 = ", noTargetIndex)

}
