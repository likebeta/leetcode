package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	bTrue, _ := isBalancedTree(root)
	return bTrue
}

func isBalancedTree(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lTrue, lLevel := isBalancedTree(root.Left)
	if !lTrue {
		return lTrue, 0
	}
	rTrue, rLevel := isBalancedTree(root.Right)
	if !rTrue {
		return rTrue, 0
	}
	if lLevel < rLevel {
		lLevel, rLevel = rLevel, lLevel
	}

	return lLevel-rLevel <= 1, lLevel + 1
}

func testOne(in string, ans bool) {
	t := helper.NewTree(in)
	t.Print(8)
	helper.Assert(isBalanced(t) == ans)
}

func main() {
	testOne("[3,9,20,null,null,15,7]", true)
	testOne("[1,2,2,3,3,null,null,4,4]", false)
	testOne("[1,null,2,null,3]", false)
}
