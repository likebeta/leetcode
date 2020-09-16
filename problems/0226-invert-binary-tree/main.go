package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Right)
	root.Right = invertTree(root.Left)
	root.Left = left
	return root
}

func main() {
	t := helper.NewTree("[4,2,7,1,3,6,9]")
	t.Print(6)
	t = invertTree(t)
	t.Print(6)
}
