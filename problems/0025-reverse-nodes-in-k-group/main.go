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
	var pre, newHead *ListNode
	newHead = head
	for head != nil {
		tail := head
		for count := 1; count < k; count++ {
			if tail.Next == nil {
				return newHead
			}
			tail = tail.Next
		}
		reverse(head, tail.Next)
		if pre == nil {
			newHead = tail
		} else {
			pre.Next = tail
		}
		pre, head = head, head.Next
	}
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	pre := b
	for a != b {
		a, a.Next, pre = a.Next, pre, a
	}
	return pre
}

func testOne(in string, k int) {
	head := helper.ParseList(in)
	head = reverseKGroup(head, k)
	helper.Log(in, "=>", head.Dump())
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
