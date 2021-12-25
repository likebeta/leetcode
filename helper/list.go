package helper

import (
	"encoding/json"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) ToArray() []int {
	curr := head
	arr := make([]int, 0)
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}

func (head *ListNode) Dump() string {
	arr := head.ToArray()
	return Dump(arr)
}

func NewList(arr []int) *ListNode {
	var pt ListNode
	p := &pt
	for i := range arr {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}
	return pt.Next
}

func ParseList(in string) *ListNode {
	var arr []int
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		log.Panic("parse failed:", err)
	}
	return NewList(arr)
}

func NewCycleList(arr []int, pos int) *ListNode {
	var pt ListNode
	var p, pPos *ListNode
	p = &pt
	for i := range arr {
		p.Next = &ListNode{Val: arr[i]}
		if pos == i {
			pPos = p.Next
		}
		p = p.Next
	}

	if pPos != nil {
		p.Next = pPos
	}
	return pt.Next
}

func ParseCycleList(in string, pos int) *ListNode {
	var arr []int
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		log.Panic("parse failed:", err)
	}
	return NewCycleList(arr, pos)
}
