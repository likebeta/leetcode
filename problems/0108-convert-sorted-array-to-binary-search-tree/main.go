package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	var ans *TreeNode
	if N := len(nums); N > 0 {
		ans = midBuild(nums, 0, N-1)
	}
	return ans
}

func midBuild(nums []int, l, r int) *TreeNode {
	var root *TreeNode
	if l <= r {
		mid := l + (r-l)/2
		root = &TreeNode{Val: nums[mid]}
		root.Left = midBuild(nums, l, mid-1)
		root.Right = midBuild(nums, mid+1, r)
	}
	return root
}

func main() {
	t := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Print(8)
}
