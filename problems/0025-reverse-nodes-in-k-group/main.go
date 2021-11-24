package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	newHead := reverse(head, tail)
	head.Next = reverseKGroup(tail, k)
	return newHead
}

// [a, b)
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	for a != b {
		a, a.Next, pre = a.Next, pre, a
	}
	return pre
}

func testOne(in string, k int) {
	head := helper.ParseList(in)
	head = reverseKGroup(head, k)
	helper.Log(in, k, "=>", head.Dump())
	helper.Print()
}

func main() {
	testOne("[]", 1)
	testOne("[1,2,3,4,5]", 1)
	testOne("[1,2,3,4,5]", 2)
	testOne("[1,2,3,4,5]", 3)
	testOne("[1,2,3,4,5]", 4)
	testOne("[1,2,3,4,5]", 5)
	testOne("[1,2,3,4,5]", 6)
}
