package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for curr != nil {
		if curr.Val > val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				break
			}
			curr = curr.Right
		}
	}
	return root
}

func testOne(in string, val int, ans string) {
	t := helper.NewTree(in)
	t = insertIntoBST(t, val)
	helper.Log(helper.InOrder(t), ans)
	t = helper.NewTree(in)
	t = insertIntoBST2(t, val)
	helper.Log(helper.InOrder(t), ans)
}

func main() {
	testOne("[4,2,7,1,3]", 5, "[1,2,3,4,5,7]")
}
