package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	p1 := lowestCommonAncestor(root.Left, p, q)
	p2 := lowestCommonAncestor(root.Right, p, q)
	if p1 == nil {
		return p2
	} else if p2 == nil {
		return p1
	} else {
		return root
	}
}

func main() {
	var t, p, q *TreeNode
	t = helper.NewTree("[3,5,1,6,2,0,8,null,null,7,4]")
	p, q = &TreeNode{Val: 5}, &TreeNode{Val: 1}
	helper.Assert(lowestCommonAncestor(t, p, q).Val == 3)

	t = helper.NewTree("[3,5,1,6,2,0,8,null,null,7,4]")
	p, q = &TreeNode{Val: 5}, &TreeNode{Val: 4}
	helper.Assert(lowestCommonAncestor(t, p, q).Val == 5)
}
