package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	var head *TreeNode
	pL, iL := len(postorder), len(inorder)
	if pL > 0 && iL > 0 && pL == iL {
		cache := make(map[int]int, iL)
		for i, v := range inorder {
			cache[v] = i
		}
		head = makeTree(cache, postorder, inorder, 0, pL-1, 0, iL-1)
	}
	return head
}

func makeTree(cache map[int]int, postorder, inorder []int, pL, pR, iL, iR int) *TreeNode {
	if pL > pR {
		return nil
	}
	head := &TreeNode{Val: postorder[pR]}
	i := cache[postorder[pR]]
	head.Left = makeTree(cache, postorder, inorder, pL, pL+i-iL-1, iL, i-1)
	head.Right = makeTree(cache, postorder, inorder, pR-(iR-i), pR-1, i+1, iR)
	return head
}

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}).Print(6)
	buildTree([]int{2, 3, 1, 4, 5}, []int{3, 2, 5, 4, 1}).Print(6)
}
