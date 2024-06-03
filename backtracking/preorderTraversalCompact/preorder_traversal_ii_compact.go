package backtracking

import (
	. "hello-algo/pkg"
)

//  在二叉树中搜索所有值为 7 的节点，请返回根节点到这些节点的路径。

func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	// 尝试
	*path = append(*path, root)
	if root.Val == 7 {
		// 记录解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

func main() {

}
