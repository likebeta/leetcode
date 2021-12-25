package main

import (
	"leetcode/helper"
	"math"
)

type TreeNode = helper.TreeNode

// 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	var midOrder func(root *TreeNode)
	min := math.MaxInt64
	var prev *TreeNode
	midOrder = func(root *TreeNode) {
		if root != nil {
			midOrder(root.Left)
			if prev != nil && root.Val - prev.Val < min {
				min = root.Val - prev.Val
			}
			prev = root
			midOrder(root.Right)
		}
	}
	midOrder(root)
	return min
}

func main() {
	t := helper.NewTree("[1,null,3,2]")
	helper.Log(getMinimumDifference(t), 1)
}
