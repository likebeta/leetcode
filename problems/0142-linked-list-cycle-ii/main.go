package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			p1 = head
			for p1 != p2 {
				p1, p2 = p1.Next, p2.Next
			}
			return p1
		}
	}
	return nil
}

func testOne(arr []int, pos int, ans int) {
	head := helper.NewCycleList(arr, pos)
	node := detectCycle(head)
	if ans == -1 {
		helper.Assert(node == nil)
	} else {
		helper.Assert(node.Val == arr[pos])
	}
}

func main() {
	testOne([]int{3, 2, 0, -4}, 1, 1)
	testOne([]int{1, 2}, 0, 0)
	testOne([]int{1}, -1, -1)
}
