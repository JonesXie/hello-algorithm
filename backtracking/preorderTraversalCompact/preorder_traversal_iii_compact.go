package backtracking

import (
	. "hello-algo/pkg"
)

// 剪枝

// 在二叉树中搜索所有值为 7 的节点，请返回根节点到这些节点的路径，并要求路径中不包含值为  3 的节点。

func preOrderIII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	// 剪枝
	if root == nil || root.Val == 3 {
		return
	}
	// 尝试
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		// 记录解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderIII(root.Left, res, path)
	preOrderIII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

func main() {

}
