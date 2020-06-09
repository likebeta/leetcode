package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var curr *ListNode
	for curr = head; n > 0; n-- {
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}
	pre := head
	for curr.Next != nil {
		pre, curr = pre.Next, curr.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func testOne(arr []int, n int) {
	head := helper.NewList(arr)
	head = removeNthFromEnd(head, n)
	helper.Log(arr, head.ToArray())
}

func main() {
	var arr []int
	arr = []int{1, 2, 3, 4, 5}
	testOne(arr, 1)
	testOne(arr, 2)
	testOne(arr, 3)
	testOne(arr, 5)
	arr = []int{1}
	testOne(arr, 1)
}
