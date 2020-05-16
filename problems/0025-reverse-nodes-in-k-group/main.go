package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	var pre, newHead *ListNode
	newHead = head
	for head != nil {
		count := 1
		start := head
		for count != k && head.Next != nil {
			head = head.Next
			count++
		}
		if count != k {
			break
		}
		reverse(start, head)
		if pre == nil {
			newHead = head
		} else {
			pre.Next = head
		}
		pre, head = start, start.Next
	}
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	if b != nil {
		pre = b.Next
	}
	for a != b {
		pre, a.Next = a.Next, pre
		pre, a = a, pre
	}
	if b != nil {
		b.Next, pre = pre, b
	}
	return pre
}

func main() {
	var list *helper.ListNode
	list = helper.NewList([]int{1, 2, 3, 4, 5})
	list = reverseKGroup(list, 2)
	helper.Log(list.ToArray())

	list = helper.NewList([]int{1, 2, 3, 4, 5})
	list = reverseKGroup(list, 3)
	helper.Log(list.ToArray())

	list = helper.NewList([]int{1, 2, 3, 4, 5})
	list = reverseKGroup(list, 5)
	helper.Log(list.ToArray())

	list = helper.NewList([]int{1, 2, 3, 4, 5})
	list = reverseKGroup(list, 6)
	helper.Log(list.ToArray())
}
