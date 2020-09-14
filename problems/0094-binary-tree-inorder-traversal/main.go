package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var (
		ans     []int
		inorder func(*TreeNode)
	)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	var ans []int

	curr := root
	for curr != nil {
		if curr.Left != nil {
			prev := curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
				continue
			}
			prev.Right = nil
		}
		ans = append(ans, curr.Val)
		curr = curr.Right
	}

	return ans
}

func testOne(in string, ans []int) {
	t := helper.NewTree(in)
	t.Print(8) // [1,3,2]
	helper.Log(inorderTraversal(t), ans)
	helper.Log(inorderTraversal2(t), ans)
}

func main() {
	testOne("[1,null,2,3]", []int{1, 3, 2})
}
