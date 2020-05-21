package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, _ := getDiameter(root)
	return ans - 1
}

func getDiameter(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	maxL, heightL := getDiameter(node.Left)
	maxR, heightR := getDiameter(node.Right)
	maxSelf := heightL + heightR + 1
	if maxR > maxSelf {
		maxSelf = maxR
	}
	if maxL > maxSelf {
		maxSelf = maxL
	}
	if heightR > heightL {
		heightL = heightR
	}
	return maxSelf, heightL + 1
}

func main() {
	t := helper.NewTree("[1,2,3,4,5]")
	helper.PrintTree(t, 6)
	helper.Assert(diameterOfBinaryTree(t) == 3)
}
