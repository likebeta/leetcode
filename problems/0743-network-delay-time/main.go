package main

import (
	"container/heap"
	"leetcode/helper"
)

// 网络延迟时间
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]*pair, n)
	for _, tuple := range times {
		from, to, weight := tuple[0]-1, tuple[1]-1, tuple[2]
		graph[from] = append(graph[from], &pair{to, weight})
	}

	costs := dijkstra(graph, n, k-1)

	var ans int
	for i := range costs {
		if costs[i] < 0 {
			return -1
		}
		ans = max(ans, costs[i])
	}
	return ans
}

func dijkstra(graph [][]*pair, n, k int) []int {
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = -1
	}
	minQ := &hp{}
	heap.Push(minQ, &pair{k, 0})
	costs[k] = 0
	for minQ.Len() > 0 {
		from := heap.Pop(minQ).(*pair)
		for _, to := range graph[from.i] {
			cost := from.w + to.w
			if costs[to.i] == -1 || costs[to.i] > cost {
				costs[to.i] = cost
				heap.Push(minQ, &pair{to.i, cost})
			}
		}

	}

	return costs
}

type pair struct {
	i int
	w int
}

type hp struct{ arr []*pair }

func (h *hp) Push(v any)         { h.arr = append(h.arr, v.(*pair)) }
func (h *hp) Pop() any           { v := h.arr[len(h.arr)-1]; h.arr = h.arr[:len(h.arr)-1]; return v }
func (h *hp) Len() int           { return len(h.arr) }
func (h *hp) Less(i, j int) bool { return h.arr[i].w < h.arr[j].w }
func (h *hp) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func testOne(in string, n, k, ans int) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(networkDelayTime(matrix, n, k) == ans)
}

func main() {
	testOne("[[2,1,1],[2,3,1],[3,4,1]]", 4, 2, 2)
	testOne("[[1,2,1]]", 2, 1, 1)
	testOne("[[1,2,1]]", 2, 2, -1)
}
