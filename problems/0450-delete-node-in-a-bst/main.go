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

		rightMin := root.Right
		for rightMin.Left != nil {
			rightMin = rightMin.Left
		}
		root.Val = rightMin.Val
		root.Right = deleteNode(root.Right, rightMin.Val)
	}
	return root
}

func testOne(in string, key int) {
	t := helper.ParseTree(in)
	t = deleteNode(t, key)
	helper.LogF("%s delete %d => %s", in, key, t.Dump())
}

func main() {
	testOne("[5,3,6,2,4,null,7]", 3)
	testOne("[5,3,6,2,4,null,7]", 0)
	testOne("[]", 0)
}
