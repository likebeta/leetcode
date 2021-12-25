package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	var inOrder func(*TreeNode, bool) int

	inOrder = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil && isLeft {
			return node.Val
		}
		return inOrder(node.Left, true) + inOrder(node.Right, false)
	}
	return inOrder(root, false)
}

func main() {
	t := helper.NewTree("[3,9,20,null,null,15,7]")
	t.Print(6)
	helper.Log(sumOfLeftLeaves(t))
}
