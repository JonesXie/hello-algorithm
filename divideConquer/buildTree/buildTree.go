package main

import (
	"fmt"
	. "hello-algo/pkg"
)

/* 构建二叉树：分治 */
func dfsBuildTree(preorder []int, inorderMap map[int]int, i, l, r int) *TreeNode {
	// 子树区间为空时终止
	if r-l < 0 {
		return nil
	}
	// 初始化根节点
	root := NewTreeNode(preorder[i])
	// 查询 m ，从而划分左右子树
	m := inorderMap[preorder[i]]
	// 子问题：构建左子树
	root.Left = dfsBuildTree(preorder, inorderMap, i+1, l, m-1)
	// 子问题：构建右子树
	root.Right = dfsBuildTree(preorder, inorderMap, i+1+m-l, m+1, r)
	// 返回根节点
	return root
}

/* 构建二叉树 */
func buildTree(preorder, inorder []int) *TreeNode {
	// 初始化哈希表，存储 inorder 元素到索引的映射
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}

	root := dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
	return root
}
func main() {
	preorder := []int{3, 9, 2, 1, 7}
	inorder := []int{9, 3, 1, 2, 7}
	fmt.Print("前序遍历 = ")
	PrintSlice(preorder)
	fmt.Print("中序遍历 = ")
	PrintSlice(inorder)

	root := buildTree(preorder, inorder)
	fmt.Println("构建的二叉树为：")
	PrintTree(root)

}
