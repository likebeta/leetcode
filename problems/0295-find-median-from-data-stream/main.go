package main

import (
	"container/heap"
	"leetcode/helper"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// MedianFinder 数据流的中位数
type MedianFinder struct {
	minQueue hp // 负数
	maxQueue hp // 正数
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	minQ, maxQ := &this.minQueue, &this.maxQueue
	if minQ.Len() == 0 {
		heap.Push(minQ, -num)
	} else if num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if minQ.Len() > maxQ.Len()+1 {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if minQ.Len() < maxQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	minQ, maxQ := &this.minQueue, &this.maxQueue
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

func testOne(action string, arg string) {
	actions := helper.ParseStrArray(action)
	args := helper.ParseFloatArray(arg)

	mf := Constructor()
	for i, a := range actions {
		if a == "addNum" {
			mf.AddNum(int(args[i]))
		} else if a == "findMedian" {
			helper.Assert(mf.FindMedian() == args[i])
		}
	}
}

func main() {
	a := `["addNum","addNum","findMedian","addNum","findMedian"]`
	b := `[1,2,1.5,3,2.0]`
	testOne(a, b)

	a = `["addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]`
	b = `[-1,-1.0,-2,-1.5,-3,-2.0,-4,-2.5,-5,-3.0]`
	testOne(a, b)
}
