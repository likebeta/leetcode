package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 根据前序和后序遍历构造二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	valToIdx := make(map[int]int, n)
	for i, v := range postorder {
		valToIdx[v] = i
	}
	return build(valToIdx, preorder, postorder, 0, n-1, 0, n-1)
}

func build(valToIdx map[int]int, preorder, postorder []int, preL, preR, postL, postR int) *TreeNode {
	if preL > preR {
		return nil
	}

	root := &TreeNode{Val: preorder[preL]}
	if preL != preR {
		// 下一个作为左节点（作为右节点也行，所以构造的树可能不唯一）
		i := valToIdx[preorder[preL+1]]
		root.Left = build(valToIdx, preorder, postorder, preL+1, preL+1+(i-postL), postL, i)
		root.Right = build(valToIdx, preorder, postorder, preR-(postR-1-i-1), preR, i+1, postR-1)
	}
	return root
}

func main() {
}
