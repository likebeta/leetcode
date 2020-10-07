package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != slow && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return fast == slow
}

func main() {
	var head *ListNode
	head = helper.ParseCycleList("[3, 2, 0, -4]", 1)
	helper.Assert(hasCycle(head) == true)

	head = helper.ParseCycleList("[1,2]", 1)
	helper.Assert(hasCycle(head) == true)

	head = helper.ParseCycleList("[1]", -1)
	helper.Assert(hasCycle(head) == false)
}
