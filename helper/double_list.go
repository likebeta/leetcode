package helper

type TDoubleListNode[T1, T2 any] struct {
	Key  T1
	Val  T2
	Prev *TDoubleListNode[T1, T2]
	Next *TDoubleListNode[T1, T2]
}

type DoubleListNode = TDoubleListNode[int, int]
