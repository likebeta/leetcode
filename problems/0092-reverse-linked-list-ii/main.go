package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		head, _ = reverseN(head, right)
	} else {
		head.Next = reverseBetween(head.Next, left-1, right-1)
	}
	return head
}

func reverseN(head *ListNode, n int) (*ListNode, *ListNode) {
	if head == nil {
		return head, nil
	} else if  n == 1 || head.Next == nil {
		return head, head.Next
	}

	newHead, last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = last
	return newHead, last
}

func testOne(in string, left, right int) {
	head := helper.ParseList(in)
	head = reverseBetween(head, left, right)
	helper.Log(in, "=>", head.Dump())
}

func main() {
	testOne("[]", 1, 1)
	testOne("[1]", 1, 1)
	testOne("[1,2]", 1, 1)
	testOne("[1,2]", 1, 2)
	testOne("[1,2]", 2, 2)
	testOne("[1,2,3]", 1, 1)
	testOne("[1,2,3]", 1, 2)
	testOne("[1,2,3]", 1, 3)
	testOne("[1,2,3]", 2, 2)
	testOne("[1,2,3]", 2, 3)
	testOne("[1,2,3]", 3, 3)
	testOne("[1,2,3]", 3, 4)
	testOne("[1,2,3]", 4, 4)
}
