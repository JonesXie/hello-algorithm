package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTee(Val int) *TreeNode {
	return &TreeNode{
		Val:   Val,
		Left:  nil,
		Right: nil,
	}
}

func initTree() *TreeNode {
	node1 := createTee(1)
	node2 := createTee(2)
	node3 := createTee(3)
	node4 := createTee(4)
	node5 := createTee(5)
	node6 := createTee(6)
	node7 := createTee(7)
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	return node1
}

func dpsTree(t *TreeNode, nums *[]int) {
	if t == nil {
		return
	}

	*nums = append(*nums, t.Val)
	dpsTree(t.Left, nums)
	dpsTree(t.Right, nums)

}

func main() {

	tree := initTree()
	nums := make([]int, 0)
	dpsTree(tree, &nums)
	fmt.Println(nums)

}
