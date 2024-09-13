package main

import (
	"leetcode/helper"
	"math"
)

type TreeNode = helper.TreeNode

// 二叉搜索子树的最大键值和
func maxSumBST(root *TreeNode) int {
	return traverse(root).maxSum
}

type bstInfo struct {
	isBST  bool
	min    int
	max    int
	sum    int
	maxSum int
}

func traverse(root *TreeNode) bstInfo {
	if root == nil {
		return bstInfo{isBST: true, min: math.MaxInt, max: math.MinInt}
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	info := bstInfo{}
	info.maxSum = max(left.maxSum, right.maxSum)
	if left.isBST && right.isBST && root.Val > left.max && root.Val < right.min {
		info.isBST = true
		info.sum = root.Val + left.sum + right.sum
		info.min = min(root.Val, left.min)
		info.max = max(root.Val, right.max)
		info.maxSum = max(info.maxSum, info.sum)
	}
	return info
}

func testOne(in string, ans int) {
	t := helper.ParseTree(in)
	helper.Assert(maxSumBST(t) == ans)
}

func main() {
	testOne("[1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]", 20)
	testOne("[4,3,null,1,2]", 2)
	testOne("[-4,-2,-5]", 0)
	testOne("[2,1,3]", 6)
	testOne("[5,4,8,3,null,6,3]", 7)
}
