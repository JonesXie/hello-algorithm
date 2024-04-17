package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// max heap 大顶堆 ---> 任意节点的值 >= 其子节点的值。
// min heap 小顶堆 ---> 任意节点的值 <= 其子节点的值。
// push()	元素入堆
// pop()	堆顶元素出堆
// peek()	访问堆顶元素（大 / 小顶堆分别为最大 / 小值）
// size()	获取堆的元素数量
// isEmpty()	判断堆是否为空

// go 在 container/heap 做了官方的实现，如果要使用堆，只需要实现下面的接口，接口中匿名嵌套了排序接口，我们也需要将排序的三个接口声明实现一遍。

// Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

type IntHeap []any

// Push heap.Interface 的方法，实现推入元素到堆
func (h *IntHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *IntHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface 的方法
func (h *IntHeap) Len() int {
	return len(*h)
}

// Less sort.Interface 的方法
func (h *IntHeap) Less(i, j int) bool {
	// 如果实现小顶堆，则需要调整为小于号
	return (*h)[i].(int) > (*h)[j].(int)
}

// Swap sort.Interface 的方法
func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top 获取堆顶元素
func (h *IntHeap) Top() any {
	return (*h)[0]
}

func main() {
	// 方法一

	h := &IntHeap{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}
	heap.Init(h)

	length := h.Len()
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", heap.Pop(h).(int))
	}

	// output:1,2,2,3,3,4,5,5,6,9,10,20,

	// 方法二
	nums := []int{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}
	h2 := &IntHeap{}

	for _, val := range nums {
		heap.Push(h2, val)
	}

	length2 := h2.Len()
	for i := 0; i < length2; i++ {
		fmt.Printf("%d,", heap.Pop(h2).(int))
	}
}
