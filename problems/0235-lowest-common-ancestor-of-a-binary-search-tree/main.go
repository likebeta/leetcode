package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}

func testOne(in string, pValue, qValue int, ans int) {
	var t, p, q *TreeNode
	t = helper.NewTree(in)
	p, q = &TreeNode{Val: pValue}, &TreeNode{Val: qValue}
	helper.Assert(lowestCommonAncestor(t, p, q).Val == ans)
	helper.Assert(lowestCommonAncestor2(t, p, q).Val == ans)
}

func main() {
	testOne("[6,2,8,0,4,7,9,null,null,3,5]", 2, 8, 6)
	testOne("[6,2,8,0,4,7,9,null,null,3,5]", 2, 4, 2)
}
