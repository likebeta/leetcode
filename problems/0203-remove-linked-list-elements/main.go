package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head != nil {
		pre, curr := head, head.Next
		for curr != nil {
			if curr.Val == val {
				pre.Next = curr.Next
			} else {
				pre = curr
			}
			curr = curr.Next
		}
	}
	return head
}

func main() {
	list := helper.NewList([]int{1, 2, 6, 3, 4, 5, 6})
	list = removeElements(list, 6)
	helper.Log(list.ToArray())
}
