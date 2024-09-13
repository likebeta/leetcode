package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	var ans []*TreeNode
	if n > 0 {
		ans = makeTrees(1, n)
	}
	return ans
}

func makeTrees(left, right int) []*TreeNode {
	var ans []*TreeNode
	if left > right {
		ans = append(ans, nil)
		return ans
	}
	for i := left; i <= right; i++ {
		lTrees := makeTrees(left, i-1)
		rTrees := makeTrees(i+1, right)
		for l := range lTrees {
			for r := range rTrees {
				head := &TreeNode{Val: i}
				head.Left = lTrees[l]
				head.Right = rTrees[r]
				ans = append(ans, head)
			}
		}

	}
	return ans
}

func main() {
	for _, t := range generateTrees(0) {
		helper.Log(t.Dump())
	}
	helper.Log("###################")
	for _, t := range generateTrees(1) {
		helper.Log(t.Dump())
	}
	helper.Log("###################")
	for _, t := range generateTrees(3) {
		helper.Log(t.Dump())
	}
}
