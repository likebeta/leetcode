package main

import (
	"leetcode/helper"
)

// 相同的树
type TreeNode = helper.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func testOne(pBfs, qBfs string, ans bool) {
	p := helper.NewTree(pBfs)
	q := helper.NewTree(qBfs)
	helper.Assert(isSameTree(p, q) == ans)
}

func main() {
	testOne("[1,2,3]", "[1,2,3]", true)
	testOne("[1,2]", "[1,null,2]", false)
	testOne("[1,2,1]", "[1,1,2]", false)
}
