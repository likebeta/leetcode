package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy ListNode
	p := &dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p, l1 = l1, l1.Next
		} else {
			p.Next = l2
			p, l2 = l2, l2.Next
		}
	}

	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}

// 合并k个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	count := len(lists)
	if count == 0 {
		return nil
	}

	l := lists[0]
	for i := 1; i < count; i++ {
		l = mergeTwoLists(l, lists[i])
	}
	return l
}

func main() {
	l1 := helper.NewList([]int{1, 2, 4})
	l2 := helper.NewList([]int{1, 3, 4})
	l3 := helper.NewList([]int{2, 4, 5})
	l := mergeKLists([]*ListNode{l1, l2, l3})
	helper.Log(l.ToArray())
}
