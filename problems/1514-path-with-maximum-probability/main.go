package main

import (
	"container/heap"
	"leetcode/helper"
)

// 概率最大的路径
func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make([][]*pair, n)
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], succProb[i]
		graph[from] = append(graph[from], &pair{to, weight})
		graph[to] = append(graph[to], &pair{from, weight})
	}

	probs := dijkstra(graph, n, start_node, end_node)
	return max(0, probs[end_node])
}

func dijkstra(graph [][]*pair, n, start, end int) []float64 {
	probs := make([]float64, n)
	for i := 0; i < n; i++ {
		probs[i] = -1
	}
	maxQ := &hp{}
	heap.Push(maxQ, &pair{start, 0})
	probs[start] = 1
	for maxQ.Len() > 0 {
		from := heap.Pop(maxQ).(*pair)
		if from.i == end {
			break
		}

		for _, to := range graph[from.i] {
			prob := probs[from.i] * to.w
			if probs[to.i] < 0 || probs[to.i] < prob {
				probs[to.i] = prob
				heap.Push(maxQ, &pair{to.i, prob})
			}
		}

	}

	return probs
}

type pair struct {
	i int
	w float64
}

type hp struct{ arr []*pair }

func (h *hp) Push(v any)         { h.arr = append(h.arr, v.(*pair)) }
func (h *hp) Pop() any           { v := h.arr[len(h.arr)-1]; h.arr = h.arr[:len(h.arr)-1]; return v }
func (h *hp) Len() int           { return len(h.arr) }
func (h *hp) Less(i, j int) bool { return h.arr[i].w > h.arr[j].w }
func (h *hp) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func testOne(sEdges, sProp string, n, start, end int, ans float64) {
	edges := helper.ParseIntMatrix(sEdges)
	succProp := helper.ParseFloatArray(sProp)
	helper.Log(maxProbability(n, edges, succProp, start, end), ans)
}

func main() {
	testOne("[[0,1],[1,2],[0,2]]", "[0.5,0.5,0.2]", 3, 0, 2, 0.25)
	testOne("[[0,1],[1,2],[0,2]]", "[0.5,0.5,0.3]", 3, 0, 2, 0.3)
	testOne("[[0,1]]", "[0.5]", 3, 0, 2, 0)
}
