package main

import "fmt"

/* 冒泡排序 */
func bubbleSort(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 优化--如果某轮“冒泡”中没有执行任何交换操作，说明数组已经完成排序，可直接返回结果

/* 冒泡排序（标志优化）*/
func bubbleSortWithFlag(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true // 记录交换元素
			}
		}
		if flag == false { // 此轮“冒泡”未交换任何元素，直接跳出
			break
		}
	}
}

func main() {
	nums := []int{4, 1, 3, 1, 5, 2}
	bubbleSort(nums)
	fmt.Println("冒泡排序完成后 nums = ", nums)

	nums1 := []int{4, 1, 3, 1, 5, 2}
	bubbleSortWithFlag(nums1)
	fmt.Println("冒泡排序完成后 nums1 = ", nums)
}
