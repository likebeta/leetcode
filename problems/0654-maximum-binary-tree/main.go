package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	maxIndex := right
	for i := left; i < right; i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	node := &TreeNode{Val: nums[maxIndex]}
	node.Left = build(nums, left, maxIndex-1)
	node.Right = build(nums, maxIndex+1, right)
	return node
}

func testOne(in string) {
	nums := helper.ParseArray(in)
	head := constructMaximumBinaryTree(nums)
	helper.Log(nums, "=>")
	helper.PrintTree(head, 4)
}

func main() {
	testOne("[]")
	testOne("[1]")
	testOne("[3,2,1,6,0,5]")
}
