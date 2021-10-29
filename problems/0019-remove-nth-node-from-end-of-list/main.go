package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

func findNthFromEnd(head *ListNode, n int) *ListNode {
	pFast := head
	for ; n > 0; n-- {
		pFast = pFast.Next
	}
	pSlow := head
	for pFast != nil {
		pSlow, pFast = pSlow.Next, pFast.Next
	}
	return pSlow
}

// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p := findNthFromEnd(dummy, n+1)
	p.Next = p.Next.Next
	return dummy.Next
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
