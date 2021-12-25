package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 路径总和 III
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return process(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func process(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	if root.Val == sum {
		count++
	}
	count += process(root.Left, sum-root.Val)
	count += process(root.Right, sum-root.Val)
	return count
}

func pathSum2(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return processWithPrefixSum(root, sum, []int{0})
}

func processWithPrefixSum(root *TreeNode, sum int, arrSum []int) int {
	if root == nil {
		return 0
	}
	sz := len(arrSum)
	arrSum = append(arrSum, arrSum[sz-1]+root.Val)
	var count int
	for i := 0; i < sz; i++ {
		if arrSum[sz]-arrSum[i] == sum {
			count++
		}
	}
	return count + processWithPrefixSum(root.Left, sum, arrSum) + processWithPrefixSum(root.Right, sum, arrSum)
}

func main() {
	t := helper.NewTree("[10,5,-3,3,2,null,11,3,-2,null,1]")
	helper.PrintTree(t, 9)
	helper.Assert(pathSum(t, 8) == 3)
	helper.Assert(pathSum2(t, 8) == 3)
}
