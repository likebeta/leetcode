package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func testOne(inA, inB string, shared string, ans int) {
	headA := helper.ParseList(inA)
	headB := helper.ParseList(inB)
	sharedHead := helper.ParseList(shared)

	headA = helper.ConcatList(headA, sharedHead)
	headB = helper.ConcatList(headB, sharedHead)

	node := getIntersectionNode(headA, headB)
	if ans == -1 {
		helper.Assert(node == nil)
	} else {
		helper.Assert(node.Val == ans)
	}
}

func main() {
	testOne("[4,1]", "[5,6,1]", "[8,4,5]", 8)
	testOne("[1,9,1]", "[3]", "[2,4]", 2)
	testOne("[2,6,4]", "[1,5]", "[]", -1)
}
