package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 恢复二叉搜索树
func recoverTree(root *TreeNode) {
	var (
		node1, node2, pre *TreeNode
		inOrder           func(node *TreeNode)
	)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if pre != nil && pre.Val > node.Val {
			node2 = node
			if node1 == nil {
				node1 = pre
			} else {
				return
			}
		}
		pre = node
		inOrder(node.Right)
	}
	inOrder(root)
	node1.Val, node2.Val = node2.Val, node1.Val
}

func recoverTree2(root *TreeNode) {
	var node1, node2, pre *TreeNode
	for root != nil {
		needCheck := false
		if root.Left == nil {
			needCheck = true
		} else {
			leftMostRight := root.Left
			for leftMostRight.Right != nil && leftMostRight.Right != root {
				leftMostRight = leftMostRight.Right
			}
			if leftMostRight.Right == root {
				needCheck = true
				leftMostRight.Right = nil
			} else {
				leftMostRight.Right = root
				root = root.Left
			}
		}
		if needCheck {
			if pre != nil && pre.Val > root.Val {
				node2 = root
				if node1 == nil {
					node1 = pre
				}
			}
			pre, root = root, root.Right
		}
	}
	node1.Val, node2.Val = node2.Val, node1.Val
}

func testOne(in string, out string) {
	inTree := helper.NewTree(in)
	recoverTree(inTree)
	helper.Assert(inTree.Dump() == out)
	inTree = helper.NewTree(in)
	recoverTree2(inTree)
	helper.Assert(inTree.Dump() == out)
}

func main() {
	testOne("[1,3,null,null,2]", "[3,1,null,null,2]")
	testOne("[3,1,4,null,null,2]", "[2,1,4,null,null,3]")
	testOne("[2,1,4,null,null,5,3]", "[2,1,4,null,null,3,5]")
	testOne("[10,5,15,0,8,13,20,2,-5,6,9,12,14,18,25]", "[10,5,15,0,8,13,20,-5,2,6,9,12,14,18,25]")
}
