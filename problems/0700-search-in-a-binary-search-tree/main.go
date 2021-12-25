package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func testOne(in string, val int) {
	t := helper.ParseTree(in)
	node := searchBST(t, val)
	helper.LogF("%s,%d => %s", in, val, node.DumpNode())
}

func main() {
	testOne("[4,2,7,1,3]", 2)
	testOne("[4,2,7,1,3]", 5)
}
