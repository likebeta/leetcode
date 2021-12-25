package main

import "leetcode/helper"

type TreeNode = helper.TreeNode

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isSymmetricTree(left.Left, right.Right) && isSymmetricTree(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var lNode, rNode *TreeNode
	qLeft, qRight := []*TreeNode{root.Left}, []*TreeNode{root.Right}
	for len(qLeft) > 0 {
		lNode, qLeft = qLeft[0], qLeft[1:]
		rNode, qRight = qRight[0], qRight[1:]
		if lNode == nil && rNode == nil {
			continue
		}
		if lNode == nil || rNode == nil || lNode.Val != rNode.Val {
			return false
		}
		qLeft = append(qLeft, lNode.Left, lNode.Right)
		qRight = append(qRight, rNode.Right, rNode.Left)
	}
	return true
}

func main() {
	var t *TreeNode
	t = helper.NewTree("[1,2,2,3,4,4,3]")
	t.Print(6)
	helper.Log(isSymmetric(t), isSymmetric2(t), true)

	t = helper.NewTree("[1,2,2,null,3,null,3]")
	t.Print(6)
	helper.Log(isSymmetric(t), isSymmetric2(t), false)
}
