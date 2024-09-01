package helper

import (
	"encoding/json"
	"log"
)

type TListNode[T any] struct {
	Val  T
	Next *TListNode[T]
}

func (node *TListNode[T]) ToArray() []T {
	curr := node
	arr := make([]T, 0)
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}

func (node *TListNode[T]) Dump() string {
	arr := node.ToArray()
	return Dump(arr)
}

type ListNode = TListNode[int]

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

func ConcatList(l1, l2 *ListNode) *ListNode {
	dummy := ListNode{Next: l1}

	head := &dummy
	for head.Next != nil {
		head = head.Next
	}
	head.Next = l2
	return dummy.Next
}
