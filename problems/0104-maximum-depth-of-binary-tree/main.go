package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	t := helper.NewTree("[3,9,20,null,null,15,7]")
	t.Print(8)
	helper.Assert(maxDepth(t) == 3)
}
