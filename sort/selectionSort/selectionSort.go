package main

import "fmt"

/* 选择排序 */
func selectionSort(nums []int) {
	n := len(nums)
	// 外循环：未排序区间为 [i, n-1]
	for i := 0; i < n-1; i++ {
		// 内循环：找到未排序区间内的最小元素
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				// 记录最小元素的索引
				k = j
			}
		}
		// 将该最小元素与未排序区间的首个元素交换
		nums[i], nums[k] = nums[k], nums[i]

	}
}

func main() {

	nums := []int{4, 1, 3, 1, 5, 2}
	selectionSort(nums)
	fmt.Println("选择排序完成后 nums = ", nums)
}
