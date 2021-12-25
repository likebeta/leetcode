package main

import (
	"leetcode/helper"
	"math"
)

type TreeNode = helper.TreeNode

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	ans := math.MaxInt32
	if root.Left != nil {
		ans = min(minDepth(root.Left), ans)
	}
	if root.Right != nil {
		ans = min(minDepth(root.Right), ans)
	}
	return ans + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func testOne(in string, ans int) {
	t := helper.NewTree(in)
	helper.Log(minDepth(t), ans)
}

func main() {
	testOne("[3,9,20,null,null,15,7]", 2)
	testOne("[1,2]", 2)
	testOne("[1]", 1)
	testOne("[1,null,2]", 2)
}
