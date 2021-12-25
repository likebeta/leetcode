package main

import (
	"fmt"
	"leetcode/helper"
	"strings"
)

type TreeNode = helper.TreeNode

// 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	var (
		ans      []string
		preOrder func(*TreeNode, []string)
	)
	preOrder = func(node *TreeNode, record []string) {
		record = append(record, fmt.Sprintf("%d", node.Val))
		if node.Left == nil && node.Right == nil {
			ans = append(ans, strings.Join(record, "->"))
			return
		}
		if node.Left != nil {
			preOrder(node.Left, record)
		}
		if node.Right != nil {
			preOrder(node.Right, record)
		}
	}
	if root != nil {
		preOrder(root, nil)
	}
	return ans
}

func testOne(in string) {
	t := helper.NewTree(in)
	t.Print(4)
	helper.Log(binaryTreePaths(t))
}

func main() {
	testOne("[1, 2, 3, null, 5]")
}
