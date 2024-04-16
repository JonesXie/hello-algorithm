package main

import (
	"container/list"
	b "hello-algo/binaryTree"
)

// 二叉树遍历

/* 层序遍历 */
func levelOrder(root *b.TreeNode) []any {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片，用于保存遍历序列
	nums := make([]any, 0)
	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*b.TreeNode)
		// 保存节点值
		nums = append(nums, node.Val)
		if node.Left != nil {
			// 左子节点入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			// 右子节点入队
			queue.PushBack(node.Right)
		}
	}
	return nums
}

/* 前序遍历 */
func preOrder(node *b.TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	// 访问优先级：根节点 -> 左子树 -> 右子树
	*nums = append(*nums, node.Val)
	preOrder(node.Left, nums)
	preOrder(node.Right, nums)
}

/* 中序遍历 */
func inOrder(node *b.TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 根节点 -> 右子树
	inOrder(node.Left, nums)
	*nums = append(*nums, node.Val)
	inOrder(node.Right, nums)
}

/* 后序遍历 */
func postOrder(node *b.TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 右子树 -> 根节点
	postOrder(node.Left, nums)
	postOrder(node.Right, nums)
	*nums = append(*nums, node.Val)
}

func main() {
	tree := b.InitTree()
	nums := make([]int, 0)

	levelOrder(tree)

	preOrder(tree, &nums)
	inOrder(tree, &nums)
	postOrder(tree, &nums)

}
