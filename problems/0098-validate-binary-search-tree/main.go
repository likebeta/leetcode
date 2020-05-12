package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

func main() {
	var node *TreeNode
	node = helper.NewTree("[2,1,3]")
	helper.PrintTree(node, 8)
	helper.Assert(isValidBST(node) == true)
	node = helper.NewTree("[5,1,4,null,null,3,6]")
	helper.PrintTree(node, 8)
	helper.Assert(isValidBST(node) == false)
}
