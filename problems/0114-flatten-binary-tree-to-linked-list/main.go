package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树展开为链表
func flatten(root *TreeNode) {
	preOrder(root)
}

func preOrder(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	left := preOrder(root.Left)
	right := preOrder(root.Right)
	if left != nil {
		left.Right, root.Right, root.Left = root.Right, root.Left, nil
	}
	if right == nil {
		return left
	} else {
		return right
	}
}

func testOne(bfs string, ans string) {
	t := helper.NewTree(bfs)
	t.Print(6)
	flatten(t)
	t.Print(6)
	helper.Assert(t.Dump() == ans)
}

func main() {
	testOne("[1,2,5,3,4,null,6]", "[1,null,2,null,3,null,4,null,5,null,6]")
	testOne("[1,2,null,3]", "[1,null,2,null,3]")
}
