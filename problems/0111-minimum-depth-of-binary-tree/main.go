package main

import (
	"leetcode/helper"
	"math"
)

type TreeNode = helper.TreeNode

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	ans := math.MaxInt32
	if root.Left != nil {
		ans = min(minDepth(root.Left), ans)
	}
	if root.Right != nil {
		ans = min(minDepth(root.Right), ans)
	}
	return ans + 1
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var depth int
	for len(queue) > 0 {
		depth++
		var tmp []*TreeNode
		for i := range len(queue) {
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}

			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = tmp
	}

	return depth
}

func testOne(in string, ans int) {
	{
		t := helper.NewTree(in)
		helper.Assert(minDepth(t) == ans)
	}

	{
		t := helper.NewTree(in)
		helper.Assert(minDepth2(t) == ans)
	}
}

func main() {
	testOne("[3,9,20,null,null,15,7]", 2)
	testOne("[1,2]", 2)
	testOne("[1]", 1)
	testOne("[1,null,2]", 2)
	testOne("[1,2,3,4,5]", 2)
}
