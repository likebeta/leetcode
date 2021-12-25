package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 打家劫舍 III
func rob(root *TreeNode) int {
	rob, noRob := doRob(root)
	return max(rob, noRob)
}

func doRob(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	robL, noRobL := doRob(node.Left)
	robR, noRobR := doRob(node.Right)
	return node.Val + noRobL + noRobR, max(robL, noRobL) + max(robR, noRobR)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testOne(in string, ans int) {
	t := helper.ParseTree(in)
	helper.Assert(rob(t) == ans)
}

func main() {
	testOne("[3,2,3,null,3,null,1]", 7)
	testOne("[3,4,5,1,3,null,1]", 9)
}
