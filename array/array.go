package main

import "fmt"

/*访问元素*/
func access(nums []int, index int) int {
	return nums[index]
}

/*插入元素*/
func insert(nums *[]int, index int, value int) {
	for i := len(*nums) - 1; i > index; i-- {
		(*nums)[i] = (*nums)[i-1]
	}
	(*nums)[index] = value
}

/*删除元素*/
func remove(nums *[]int, index int) {
	for i := index; i < len(*nums)-1; i++ {
		(*nums)[i] = (*nums)[i+1]
	}
}

func main() {

	/* 初始化数组 */
	// var arr [5]int

	// 在 Go 中，指定长度时（[5]int）为数组，不指定长度时（[]int）为切片
	// 由于 Go 的数组被设计为在编译期确定长度，因此只能使用常量来指定长度
	// 为了方便实现扩容 extend() 方法，以下将切片（Slice）看作数组（Array）
	nums := []int{1, 2, 3, 4, 5}

	/*访问元素*/
	fmt.Println(access(nums, 0))

	/*插入元素*/
	insert(&nums, 1, 6)
	fmt.Println(nums)

	/*删除元素*/
	remove(&nums, 3)
	fmt.Println(nums)

}
