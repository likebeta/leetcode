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

func testOne(in string, pos int, ans bool) {
	head := helper.ParseCycleList(in, pos)
	helper.Assert(hasCycle(head) == ans)
}

func main() {
	testOne("[3, 2, 0, -4]", 1, true)
	testOne("[1,2]", 1, true)
	testOne("[1]", -1, false)
}
