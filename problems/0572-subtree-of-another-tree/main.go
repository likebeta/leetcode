package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 另一个树的子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSameTree(s, t) {
		return true
	}
	return s != nil && (isSubtree(s.Left, t) || isSubtree(s.Right, t))
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

func main() {
	var t1, t2 *TreeNode
	t1 = helper.NewTree("[3,4,5,1,2]")
	t2 = helper.NewTree("[4,1,2]")
	helper.Assert(isSubtree(t1, t2) == true)

	t1 = helper.NewTree("[3,4,5,1,2,null,null,null,null,0]")
	t2 = helper.NewTree("[4,1,2]")
	helper.Assert(isSubtree(t1, t2) == false)
}
