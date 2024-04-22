package main

import (
	"fmt"
	"testing"
)

// question:给定一个长度为 n 的数组 nums ，元素按从小到大的顺序排列且不重复。请查找并返回元素 target 在该数组中的索引。若数组不包含该元素，则返回-1

// 二分查找(双闭区间)
func binarySearch(nums []int, target int) int {
	// 初始化双闭区间 [0, n-1] ，即 i, j 分别指向数组首元素、尾元素
	i := 0
	j := len(nums) - 1
	// 循环，当搜索区间为空时跳出（当 i > j 时为空）
	for i <= j {
		// i 和 j 都是 int 类型，因此 i+j 可能会超出 int 类型的取值范围。为了避免大数越界，我们通常采用公式 i+ (j-i)/2 来计算中点。
		mid := i + (j-i)/2 // 计算中点索引 m
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 此情况说明 target 在区间 [m+1, j] 中
			i = mid + 1
		} else { // 此情况说明 target 在区间 [i, m-1] 中
			j = mid - 1
		}
	}

	return -1
}

/* 二分查找（左闭右开区间） */
func binarySearchLCRO(nums []int, target int) int {
	// 初始化左闭右开区间 [0, n) ，即 i, j 分别指向数组首元素、尾元素+1
	i, j := 0, len(nums)
	// 循环，当搜索区间为空时跳出（当 i = j 时为空）
	for i < j {
		m := i + (j-i)/2      // 计算中点索引 m
		if nums[m] < target { // 此情况说明 target 在区间 [m+1, j) 中
			i = m + 1
		} else if nums[m] > target { // 此情况说明 target 在区间 [i, m) 中
			j = m
		} else { // 找到目标元素，返回其索引
			return m
		}
	}
	// 未找到目标元素，返回 -1
	return -1
}

//  测试-二分查找

func TestBinarySearch(t *testing.T) {
	var (
		target   = 6
		nums     = []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
		expected = 2
	)
	// 在数组中执行二分查找
	actual := binarySearch(nums, target)
	fmt.Println("目标元素 6 的索引 =", actual)
	if actual != expected {
		t.Errorf("目标元素 6 的索引 = %d, 应该为 %d", actual, expected)
	}
}

/*二分查找插入点*/

// Q:给定一个长度为  n 的有序数组 nums 和一个元素 target ，数组不存在重复元素。现将 target 插入数组 nums 中，并保持其有序性。若数组中已存在元素 target ，则插入到其左方。请返回插入后 target 在数组中的索引。

// 二分查找插入点（无重复元素）
func binarySearchInsertionSimple(nums []int, target int) int {

	i, j := 0, len(nums)-1
	// 循环，当搜索区间为空时跳出（当 i > j 时为空）
	for i <= j {

		mid := i + (j-i)/2 // 计算中点索引 m
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 此情况说明 target 在区间 [m+1, j] 中
			i = mid + 1
		} else { // 此情况说明 target 在区间 [i, m-1] 中
			j = mid - 1
		}
	}

	return -1
}

// Q: 在上题基础上，规定数组可能包含重复元素

/* 二分查找插入点（存在重复元素） */
func binarySearchInsertion(nums []int, target int) int {
	// 初始化双闭区间 [0, n-1]
	i, j := 0, len(nums)-1
	for i <= j {
		// 计算中点索引 m
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在区间 [i, m-1] 中
			j = m - 1
		} else {
			// 首个小于 target 的元素在区间 [i, m-1] 中
			j = m - 1
		}
	}
	// 返回插入点 i
	return i
}

// 测试-插入

func TestBinarySearchInsertion(t *testing.T) {
	// 无重复元素的数组
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	fmt.Println("数组 nums =", nums)

	// 二分查找插入点
	for _, target := range []int{6, 9} {
		index := binarySearchInsertionSimple(nums, target)
		fmt.Println("元素", target, "的插入点的索引为", index)
	}

	// 包含重复元素的数组
	nums = []int{1, 3, 6, 6, 6, 6, 6, 10, 12, 15}
	fmt.Println("\n数组 nums =", nums)

	// 二分查找插入点
	for _, target := range []int{2, 6, 20} {
		index := binarySearchInsertion(nums, target)
		fmt.Println("元素", target, "的插入点的索引为", index)
	}
}

// Q:给定一个长度为 n 的有序数组 nums ，其中可能包含重复元素。请返回数组中最左一个元素 target 的索引。若数组中不包含该元素，则返回  -1 。

/* 二分查找最左一个 target */
func binarySearchLeftEdge(nums []int, target int) int {
	// 等价于查找 target 的插入点
	i := binarySearchInsertion(nums, target)
	// 未找到 target ，返回 -1
	if i == len(nums) || nums[i] != target {
		return -1
	}
	// 找到 target ，返回索引 i
	return i
}

/* 二分查找最右一个 target */
func binarySearchRightEdge(nums []int, target int) int {
	// 转化为查找最左一个 target + 1
	i := binarySearchInsertion(nums, target+1)
	// j 指向最右一个 target ，i 指向首个大于 target 的元素
	j := i - 1
	// 未找到 target ，返回 -1
	if j == -1 || nums[j] != target {
		return -1
	}
	// 找到 target ，返回索引 j
	return j
}

// 测试-边界

func TestBinarySearchEdge(t *testing.T) {
	// 包含重复元素的数组
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	fmt.Println("\n数组 nums = ", nums)

	// 二分查找左边界和右边界
	for _, target := range []int{6, 7} {
		index := binarySearchLeftEdge(nums, target)
		fmt.Println("最左一个元素", target, "的索引为", index)

		index = binarySearchRightEdge(nums, target)
		fmt.Println("最右一个元素", target, "的索引为", index)
	}
}

func main() {

}
