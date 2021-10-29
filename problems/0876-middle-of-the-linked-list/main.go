package main

import (
	"leetcode/helper"
)

type ListNode = helper.ListNode

// 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	pSlow, pFast := head, head
	for pFast != nil && pFast.Next != nil {
		pSlow, pFast = pSlow.Next, pFast.Next.Next
	}
	return pSlow
}

func testOne(in string) {
	head := helper.ParseList(in)
	node := middleNode(head)
	if node == nil {
		helper.Log(in, "=>")
	} else {
		helper.Log(in, "=>", node.Val)
	}
}

func main()  {
	testOne("[]")
	testOne("[1]")
	testOne("[1,2]")
	testOne("[1,2,3]")
	testOne("[1,2,3,4]")
	testOne("[1,2,3,4,5]")
}
