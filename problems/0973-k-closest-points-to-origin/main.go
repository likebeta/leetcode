package main

import (
	"container/heap"
	"leetcode/helper"
)

// 最接近原点的 K 个点
type pair struct {
	dist  int
	point []int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosest(points [][]int, k int) (ans [][]int) {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h) // O(k) 初始化堆
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}

func testOne(in string, k int, ans ...string) bool {
	var points [][]int
	helper.Load([]byte(in), &points)
	result := kClosest(points, k)
	s := helper.Dump(result)
	for _, v := range ans {
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	helper.Assert(testOne("[[1,3],[-2,2]]", 1, "[[-2,2]]"))
	helper.Assert(testOne("[[3,3],[5,-1],[-2,4]]", 2, "[[3,3],[-2,4]]", "[[-2,4],[3,3]]"))
}
