package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode
type TreeNode = helper.TreeNode

// 有序链表转换二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	var (
		getMedian func(*ListNode, *ListNode) *ListNode
		buildTree func(*ListNode, *ListNode) *TreeNode
	)

	getMedian = func(left *ListNode, right *ListNode) *ListNode {
		fast, slow := left, left
		for fast != right && fast.Next != right {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	buildTree = func(left *ListNode, right *ListNode) *TreeNode {
		if left == right {
			return nil
		}
		mid := getMedian(left, right)
		root := &TreeNode{Val: mid.Val}
		root.Left = buildTree(left, mid)
		root.Right = buildTree(mid.Next, right)
		return root
	}

	return buildTree(head, nil)
}

func sortedListToBST2(head *ListNode) *TreeNode {
	var (
		globalHead = head
		getLength  func(*ListNode) int
		buildTree  func(int, int) *TreeNode
	)
	getLength = func(head *ListNode) int {
		length := 0
		for curr := head; curr != nil; curr = curr.Next {
			length++
		}
		return length
	}

	buildTree = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)/2
		root := &TreeNode{}
		root.Left = buildTree(left, mid-1)
		root.Val = globalHead.Val
		globalHead = globalHead.Next
		root.Right = buildTree(mid+1, right)
		return root
	}
	return buildTree(0, getLength(head)-1)
}

func testOne(in string, ans string) {
	list := helper.ParseList(in)
	t1 := sortedListToBST(list)
	helper.Assert(t1.Dump() == ans)
	t2 := sortedListToBST2(list)
	helper.Log(t2.Dump(), ans)
}

func main() {
	testOne("[-10,-3,0,5,9]", "[0,-3,9,-10,null,5]")
}
