package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 求根到叶子节点数字之和
func sumNumbers(root *TreeNode) int {
	var (
		midOrder func(int, *TreeNode) int
	)
	midOrder = func(pre int, node *TreeNode) int {
		if node == nil {
			return 0
		}
		pre = pre*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return pre
		}
		return midOrder(pre, node.Left) + midOrder(pre, node.Right)
	}

	return midOrder(0, root)
}

func testOne(in string, ans int) {
	arr := helper.ParseTree(in)
	helper.Assert(sumNumbers(arr) == ans)
}

func main() {
	testOne("[]", 0)
	testOne("[1]", 1)
	testOne("[1,2]", 12)
	testOne("[1,null,3]", 13)
	testOne("[1,2,3]", 25)
	testOne("[4,9,0,5,1]", 1026)
}
