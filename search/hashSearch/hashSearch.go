package main

import "fmt"

// 给定一个整数数组 nums 和一个目标元素 target ，请在数组中搜索“和”为 target 的两个元素，并返回它们的数组索引。返回任意一个解即可。

// 方法一：线性查找，时间换空间

func twoSumBruteForce(nums []int, target int) []int {
	size := len(nums)
	// 两层循环，时间复杂度为 O(n^2)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：哈希查找，空间换时间

func twoSumHashTable(nums []int, target int) []int {
	hashTable := map[int]int{}

	for i, v := range nums {
		if preVal, ok := hashTable[target-v]; ok {
			return []int{preVal, i}
		}
		hashTable[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 13
	site := twoSumBruteForce(nums, target)
	fmt.Println(site)
	site2 := twoSumHashTable(nums, target)
	fmt.Println(site2)
}
