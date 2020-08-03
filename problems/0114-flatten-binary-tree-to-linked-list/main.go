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

func flatten2(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
			continue
		}
		leftMostRight := curr.Left
		for leftMostRight.Right != nil && leftMostRight.Right != curr {
			leftMostRight = leftMostRight.Right
		}
		if leftMostRight.Right == curr {
			leftMostRight.Right, curr.Right, curr.Left = curr.Right, curr.Left, nil
			curr = leftMostRight.Right
		} else {
			leftMostRight.Right = curr
			curr = curr.Left
		}
	}
}

func flatten3(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			leftMostRight := next
			for leftMostRight.Right != nil {
				leftMostRight = leftMostRight.Right
			}
			leftMostRight.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func testOne(bfs string, ans string) {
	{
		t := helper.NewTree(bfs)
		t.Print(6)
		flatten(t)
		t.Print(6)
		helper.Assert(t.Dump() == ans)
	}

	{
		t := helper.NewTree(bfs)
		t.Print(6)
		flatten2(t)
		t.Print(6)
		helper.Log(t.Dump() , ans)
	}

	{
		t := helper.NewTree(bfs)
		t.Print(6)
		flatten3(t)
		t.Print(6)
		helper.Log(t.Dump() , ans)
	}
}

func main() {
	testOne("[1,2,5,3,4,null,6]", "[1,null,2,null,3,null,4,null,5,null,6]")
	testOne("[1,2,null,3]", "[1,null,2,null,3]")
}
