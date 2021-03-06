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

func isPalindrome2(head *ListNode) bool {
	var f func(*ListNode) bool
	f = func(node *ListNode) bool {
		if node != nil {
			if !f(node.Next) || node.Val != head.Val {
				return false
			}
			head = head.Next
		}
		return true
	}
	return f(head)
}

func testOne(arr []int, ans bool) {
	{
		list := helper.NewList(arr)
		helper.Assert(isPalindrome(list) == ans)
		helper.Assert(helper.DumpArray(arr) == list.Dump())
	}
	{
		list := helper.NewList(arr)
		helper.Assert(isPalindrome2(list) == ans)
	}
}

func main() {
	testOne(nil, true)
	testOne([]int{}, true)
	testOne([]int{1}, true)
	testOne([]int{1, 2}, false)
	testOne([]int{1, 1}, true)
	testOne([]int{1, 1, 2}, false)
	testOne([]int{1, 2, 1}, true)
	testOne([]int{1, 2, 2, 1}, true)
	testOne([]int{1, 2, 2, 3}, false)
	testOne([]int{1, 2, 2, 2, 1}, true)
}
