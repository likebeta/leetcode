package helper

import (
	"fmt"
	"strings"
)

type TLinkedListNode[T1, T2 any] struct {
	Key  T1
	Val  T2
	Prev *TLinkedListNode[T1, T2]
	Next *TLinkedListNode[T1, T2]
}

type TLinkedList[T1, T2 any] struct {
	head   *TLinkedListNode[T1, T2]
	tail   *TLinkedListNode[T1, T2]
	length int
}

func (l *TLinkedList[T1, T2]) Remove(node *TLinkedListNode[T1, T2]) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	l.length--
}

func (l *TLinkedList[T1, T2]) RemoveLast() *TLinkedListNode[T1, T2] {
	node := l.tail.Prev
	l.Remove(node)
	return node
}

func (l *TLinkedList[T1, T2]) InsertFirst(node *TLinkedListNode[T1, T2]) {
	node.Prev, node.Next = l.head, l.head.Next
	l.head.Next, l.head.Next.Prev = node, node

	l.length++
}

func (l *TLinkedList[T1, T2]) Empty() bool {
	return l.length == 0
}

func (l *TLinkedList[T1, T2]) Len() int {
	return l.length
}

func (l *TLinkedList[T1, T2]) Dump() (string, string) {
	var lToR, rToL string
	{
		var lines []string
		for node := l.head; node != nil; node = node.Next {
			lines = append(lines, fmt.Sprintf("[%d,%d]", node.Key, node.Val))
		}
		lToR = strings.Join(lines, "=>")
	}

	{
		var lines []string
		for node := l.tail; node != nil; node = node.Prev {
			lines = append(lines, fmt.Sprintf("[%d,%d]", node.Key, node.Val))
		}
		rToL = strings.Join(lines, "<=")
	}
	return lToR, rToL
}

type LinkedListNode = TLinkedListNode[int, int]
type LinkedList = TLinkedList[int, int]

func NewLinkedList() *LinkedList {
	var head, tail LinkedListNode
	head.Next, tail.Prev = &tail, &head

	return &LinkedList{head: &head, tail: &tail}
}
