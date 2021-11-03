package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var midOrder func(*TreeNode)
	midOrder = func(node *TreeNode) {
		if node == nil || k < 0 {
			return
		}
		midOrder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		midOrder(node.Right)
	}
	midOrder(root)
	return ans
}

func testOne(in string, k int) {
	t := helper.ParseTree(in)
	v := kthSmallest(t, k)
	helper.Log(in, "=>", v)
}

func main() {
	testOne("[3,1,4,null,2]", 1)
	testOne("[5,3,6,2,4,null,null,1]", 3)
}
