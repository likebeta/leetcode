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

func testOne(in string) {
	head := helper.ParseList(in)
	head = reverseList(head)
	helper.Log(in, "=>", head.Dump())
}

func main() {
	testOne("[]")
	testOne("[1]")
	testOne("[1,2]")
	testOne("[1,2,3]")
}
