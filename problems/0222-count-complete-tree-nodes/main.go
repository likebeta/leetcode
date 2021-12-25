package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	} else {
		return countNodes(root.Left) + (1 << right)
	}
}

func countLevel(node *TreeNode) int {
	var level int
	for node != nil {
		level++
		node = node.Left
	}
	return level
}

func testOne(in string, ans int) {
	t := helper.ParseTree(in)
	helper.Assert(countNodes(t) == ans)
}

func main() {
	testOne("[1,2,3,4,5,6]", 6)
}
