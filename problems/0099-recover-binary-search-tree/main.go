package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 恢复二叉搜索树
func recoverTree(root *TreeNode)  {
	var (
		node1, node2, pre *TreeNode
		inOrder func(node *TreeNode)
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

func testOne(in string, out string) {
	inTree := helper.NewTree(in)
	recoverTree(inTree)
	helper.Assert(inTree.Dump() == out)
}

func main() {
	testOne("[1,3,null,null,2]", "[3,1,null,null,2]")
	testOne("[3,1,4,null,null,2]", "[2,1,4,null,null,3]")
	testOne("[2,1,4,null,null,5,3]", "[2,1,4,null,null,3,5]")
}
