package backtracking

import (
	. "hello-algo/pkg"
)

//给定一棵二叉树，搜索并记录所有值为 7 的节点，请返回节点列表。

func preOrderI(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Val == 7 {
		*res = append(*res, root)
	}
	preOrderI(root.Left, res)
	preOrderI(root.Right, res)
}

func main() {

}
