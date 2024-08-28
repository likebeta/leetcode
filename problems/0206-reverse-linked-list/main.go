package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}

func testOne(in string, out string, f func(*ListNode) *ListNode) {
	head := helper.ParseList(in)
	head = f(head)
	helper.Assert(head.Dump() == out)
}

func main() {
	testOne("[]", "[]", reverseList)
	testOne("[1]", "[1]", reverseList)
	testOne("[1,2]", "[2,1]", reverseList)
	testOne("[1,2,3]", "[3,2,1]", reverseList)

	testOne("[]", "[]", reverseList2)
	testOne("[1]", "[1]", reverseList2)
	testOne("[1,2]", "[2,1]", reverseList2)
	testOne("[1,2,3]", "[3,2,1]", reverseList2)
}
