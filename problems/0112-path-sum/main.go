package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 路径总和
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func main() {
	var node *TreeNode
	node = helper.NewTree("[5,4,8,11,null,13,4,7,2,null,null,null,1]")
	helper.PrintTree(node, 8)
	helper.Assert(hasPathSum(node, 22) == true)
}
