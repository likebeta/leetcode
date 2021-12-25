package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	var (
		ans      []int
		preOrder func(root *TreeNode)
	)

	preOrder = func(root *TreeNode) {
		if root != nil {
			ans = append(ans, root.Val)
			preOrder(root.Left)
			preOrder(root.Right)
		}
	}
	preOrder(root)
	return ans
}

func preorderTraversal2(root *TreeNode) (vals []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func preorderTraversal3(root *TreeNode) []int {
	var ans []int
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				ans = append(ans, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			ans = append(ans, p1.Val)
		}
		p1 = p1.Right
	}
	return ans
}

func testOne(in string, ans string) {
	type preOrder func(root *TreeNode) []int
	funcArr := []preOrder{preorderTraversal, preorderTraversal2, preorderTraversal3}
	for _, func_ := range funcArr {
		t := helper.NewTree(in)
		result := func_(t)
		helper.Assert(helper.DumpArray(result) == ans)
	}
}

func main() {
	testOne("[1,null,2,3]", "[1,2,3]")
}
