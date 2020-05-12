package helper

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) ToArray() []int {
	curr := head
	var arr []int
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
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
