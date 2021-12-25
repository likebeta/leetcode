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

	fast = reverse(slow.Next)
	left, right := head, fast
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}
	slow.Next = reverse(fast)
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head.Next = head.Next, pre
		pre, head = head, pre
	}
	return pre
}

// 递归方法
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

func testOne(in string, ans bool) {
	{
		list := helper.ParseList(in)
		helper.Assert(isPalindrome(list) == ans)
	}
	{
		list := helper.ParseList(in)
		helper.Assert(isPalindrome2(list) == ans)
	}
}

func main() {
	testOne("[]", true)
	testOne("[1]", true)
	testOne("[1,2]", false)
	testOne("[1,1]", true)
	testOne("[1,1,2]", false)
	testOne("[1,2,1]", true)
	testOne("[1,2,2,1]", true)
	testOne("[1,2,2,3]", false)
	testOne("[1,2,2,2,1]", true)
}
