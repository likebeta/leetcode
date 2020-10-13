package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil && curr.Next != nil {
		if prev == nil {
			head = curr.Next
		} else {
			prev.Next = curr.Next
		}
		curr.Next, curr.Next.Next = curr.Next.Next, curr
		prev, curr = curr, curr.Next
	}
	return head
}

func testOne(arr []int, in string) {
	list := helper.NewList(arr)
	list = swapPairs(list)
	helper.Log(list.Dump(), in)
}

func main() {
	testOne(nil, "[]")
	testOne([]int{}, "[]")
	testOne([]int{1}, "[1]")
	testOne([]int{1, 2}, "[2,1]")
	testOne([]int{1, 2, 3}, "[2,1,3]]")
	testOne([]int{1, 2, 3, 4}, "[2,1,4,3]")
	testOne([]int{1, 2, 3, 4, 5}, "[2,1,4,3,5]")
}
