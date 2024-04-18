package main

import (
	"container/heap"
	"fmt"
)

// question:给定一个长度为 n 的无序数组 nums ，请返回数组中最大的 k 个元素。

type minHeap []any

func (h *minHeap) Push(x any) {
	*h = append(*h, x)
}

func (h *minHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].(int) < (*h)[j].(int)
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Top() any {
	return (*h)[0]
}

/* 基于堆查找数组中最大的 k 个元素 */
func topKHeap(nums []int, k int) *minHeap {
	// 初始化小顶堆
	h := &minHeap{}
	heap.Init(h)
	// 将数组的前 k 个元素入堆
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	// 从第 k+1 个元素开始，保持堆的长度为 k
	for i := k; i < len(nums); i++ {
		// 若当前元素大于堆顶元素，则将堆顶元素出堆、当前元素入堆
		if nums[i] > h.Top().(int) {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return h
}

func main() {

	kVal := topKHeap([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4)
	fmt.Print(kVal)

}
