package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func main() {
	t1 := helper.NewTree("[1, 3, 2, 5]")
	t2 := helper.NewTree("[2, 1, 3, null, 4, null, 7]")
	helper.Log(mergeTrees(t1, t2).Dump(), "[3,4,5,5,4,null,7]")
}
