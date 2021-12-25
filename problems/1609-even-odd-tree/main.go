package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 奇偶树
func isEvenOddTree(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	for level := 0; len(nodes) > 0; level++ {
		var queue []*TreeNode
		for i, node := range nodes {
			if node.Val%2 == level%2 {
				return false
			}
			if i > 0 && ((level%2 == 0 && node.Val <= nodes[i-1].Val) ||
				(level%2 == 1 && node.Val >= nodes[i-1].Val)) {
				return false
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		nodes = queue
	}
	return true
}

func testOne(in string, ans bool) {
	root := helper.ParseTree(in)
	helper.Assert(isEvenOddTree(root) == ans)
}

func main() {
	testOne("[1,10,4,3,null,7,9,12,8,6,null,null,2]", true)
	testOne("[5,4,2,3,3,7]", false)
	testOne("[5,9,1,3,5,7]", false)
	testOne("[1]", true)
	testOne("[11,8,6,1,3,9,11,30,20,18,16,12,10,4,2,17]", true)
}
