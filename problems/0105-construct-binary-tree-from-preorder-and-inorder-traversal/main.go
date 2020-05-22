package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var head *TreeNode
	pL, iL := len(preorder), len(inorder)
	if pL > 0 && iL > 0 && pL == iL {
		cache := make(map[int]int, iL)
		for i, v := range inorder {
			cache[v] = i
		}
		head = makeTree(cache, preorder, inorder, 0, pL-1, 0, iL-1)
	}
	return head
}

func makeTree(cache map[int]int, preorder, inorder []int, pL, pR, iL, iR int) *TreeNode {
	if pL > pR {
		return nil
	}
	head := &TreeNode{Val: preorder[pL]}
	i := cache[preorder[pL]]
	head.Left = makeTree(cache, preorder, inorder, pL+1, pL+i-iL, iL, i-1)
	head.Right = makeTree(cache, preorder, inorder, pR-(iR-i)+1, pR, i+1, iR)
	return head
}

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).Print(5)
	buildTree([]int{1, 2, 3, 4, 5}, []int{2, 3, 1, 4, 5}).Print(5)
}
