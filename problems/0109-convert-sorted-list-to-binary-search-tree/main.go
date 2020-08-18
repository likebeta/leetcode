package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode
type TreeNode = helper.TreeNode

// 有序链表转换二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

func testOne(in []int) {
	list := helper.NewList(in)
	t := sortedListToBST(list)
	t.Print(8)
}

func main() {
	testOne([]int{-10, -3, 0, 5, 9})
}
