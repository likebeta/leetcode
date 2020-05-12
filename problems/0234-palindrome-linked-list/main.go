package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// a->b->c->d     a->b->     c<-d
// a->b->c->d->e  a->b->c->  d<-e

// 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	ans := true
	fast = reverse(slow.Next)
	left, right := head, fast
	for left != nil && right != nil {
		if left.Val != right.Val {
			ans = false
			break
		}
		left, right = left.Next, right.Next
	}
	slow.Next = reverse(fast)
	return ans
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head.Next = head.Next, pre
		pre, head = head, pre
	}
	return pre
}

func main() {
	testOne := func(arr []int) {
		list := helper.NewList(arr)
		result := isPalindrome(list)
		helper.Log(arr, result, list.ToArray())
	}
	testOne(nil)
	testOne([]int{})
	testOne([]int{1})
	testOne([]int{1, 2})
	testOne([]int{1, 1})
	testOne([]int{1, 1, 2})
	testOne([]int{1, 2, 1})
	testOne([]int{1, 2, 2, 1})
	testOne([]int{1, 2, 2, 3})
	testOne([]int{1, 2, 2, 2, 1})
}
