package main

import "leetcode/helper"

type ListNode = helper.ListNode

// 移除重复节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	occurred := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}

func testOne(arr []int, ans []int) {
	head := helper.NewList(arr)
	head = removeDuplicateNodes(head)
	helper.Log(arr, head.ToArray(), ans)
}

func main() {
	testOne([]int{1, 2, 3, 3, 2, 1}, []int{1, 2, 3})
	testOne([]int{1, 1, 1, 1, 2}, []int{1, 2})
}
