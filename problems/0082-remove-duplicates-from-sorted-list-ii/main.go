package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 删除排序链表中的重复元素 II
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			pre, pre.Next = head, head
			head = head.Next
		} else {
			x := head.Val
			for head != nil && head.Val == x {
				head = head.Next
			}
			pre.Next = head
		}
	}
	return dummy.Next
}

func testOne(in string, out string) {
	list := helper.ParseList(in)
	list = deleteDuplicates(list)
	helper.Log(in, "=>", list.Dump(), out)
}

func main() {
	testOne("[]", "[]")
	testOne("[1]", "[1]")
	testOne("[1,1]", "[]")
	testOne("[1,1,2,2]", "[]")
	testOne("[1,2,3,3,4,4,5]", "[1,2,5]")
	testOne("[1,1,1,2,3]", "[2,3]")
	testOne("[-1,0,0,0,0,3,3]", "[-1]")
}
