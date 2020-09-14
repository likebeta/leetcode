package main

import (
	"fmt"
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

func main() {
	t := helper.NewTree("[1,null,2,3]")
	t.Print(8)
	fmt.Println(inorderTraversal(t))  // [1,3,2]
}
