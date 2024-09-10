package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		switchWithRightMin(root)
		// switchWithLeftMax(root)
	}
	return root
}

func switchWithRightMin(root *TreeNode) {
	rightMin := root.Right
	for rightMin.Left != nil {
		rightMin = rightMin.Left
	}
	root.Val = rightMin.Val
	root.Right = deleteNode(root.Right, rightMin.Val)
}

func switchWithLeftMax(root *TreeNode) {
	leftMax := root.Left
	for leftMax.Right != nil {
		leftMax = leftMax.Right
	}
	root.Val = leftMax.Val
	root.Left = deleteNode(root.Left, leftMax.Val)
}

func testOne(in string, key int, ans string, others ...string) {
	t := helper.ParseTree(in)
	t = deleteNode(t, key)
	s := t.Dump()
	helper.LogF("%s delete %d => %s", in, key, s)

	yes := ans == s
	for i := 0; !yes && i < len(others); i++ {
		if others[i] == s {
			yes = true
		}
	}
	helper.Assert(yes)
}

func main() {
	testOne("[5,3,6,2,4,null,7]", 3, "[5,2,6,null,4,null,7]", "[5,4,6,2,null,null,7]")
	testOne("[5,3,6,2,4,null,7]", 0, "[5,3,6,2,4,null,7]")
	testOne("[]", 0, "[]")
}
