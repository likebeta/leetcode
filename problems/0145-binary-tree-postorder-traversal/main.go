package main

import (
	"fmt"
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	var (
		ans       []int
		postOrder func(root *TreeNode)
	)

	postOrder = func(root *TreeNode) {
		if root != nil {
			postOrder(root.Left)
			postOrder(root.Right)
			ans = append(ans, root.Val)
		}
	}
	postOrder(root)
	return ans
}

func postorderTraversal2(root *TreeNode) []int {
	var ans []int

	if root != nil {
		var prev *TreeNode
		stack := []*TreeNode{root}
		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			if curr.Left != nil && curr.Left != prev && (prev == nil || curr.Right != prev) {
				for curr.Left != nil {
					stack = append(stack, curr.Left)
					curr = curr.Left
				}
			} else if curr.Right != nil && (prev == nil || curr.Right != prev) {
				stack = append(stack, curr.Right)
			} else {
				ans = append(ans, curr.Val)
				prev = curr
				stack = stack[:len(stack)-1]
			}
		}
	}
	return ans
}

func printStack(stack []*TreeNode, prev *TreeNode) {
	fmt.Printf("%d  == %v", len(stack), prev)
	for _, t := range stack {
		fmt.Print(" ", t.Val)
	}
	fmt.Println()
}

func testOne(in string, ans string) {
	t := helper.NewTree(in)
	helper.Log(postorderTraversal(t), ans)
	helper.Log(postorderTraversal2(t), ans)
}

func main() {
	testOne("[1,null,2,3]", "[3,2,1]")
	testOne("[3,1,2]", "[1,2,3]")
}
