package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		if pre == dummy || head.Val != pre.Val {
			pre, pre.Next = head, head
		}
		head = head.Next
	}
	pre.Next = nil
	return dummy.Next
}

func testOne(in string, out string) {
	list := helper.ParseList(in)
	list = deleteDuplicates(list)
	helper.Log(in, "=>", list.Dump(), out)
}

func main() {
	testOne("[]", "[]")
	testOne("[1,2,3,3,4,4,5]", "[1,2,3,4,5]")
	testOne("[1,1,1,2,3]", "[1,2,3]")
	testOne("[-1,0,0,0,0,3,3]", "[-1,0,3]")
}
