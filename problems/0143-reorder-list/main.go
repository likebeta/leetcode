package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 重排链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(node *ListNode) *ListNode {
	slow, fast := node, node
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(node *ListNode) *ListNode {
	var pre *ListNode
	for node != nil {
		pre, node, node.Next = node, node.Next, pre
	}
	return pre
}

func mergeList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1, l1.Next, l2, l2.Next = l1.Next, l2, l2.Next, l1.Next
	}
}

func testOne(in string, ans string) {
	head := helper.ParseList(in)
	reorderList(head)
	helper.Assert(head.Dump() == ans)
}

func main() {
	testOne("[]", "[]")
	testOne("[1]", "[1]")
	testOne("[1,2]", "[1,2]")
	testOne("[1,2,3]", "[1,3,2]")
	testOne("[1,2,3,4]", "[1,4,2,3]")
	testOne("[1,2,3,4,5]", "[1,5,2,4,3]")
	testOne("[1,2,3,4,5,6]", "[1,6,2,5,3,4]")
}
