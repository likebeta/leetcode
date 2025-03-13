package main

import (
	"container/heap"
	"leetcode/helper"
	"sort"
)

// 最低成本连通所有城市
func minimumCost(n int, connections [][]int) int {
	return kruskal(n, connections)
}

func kruskal(n int, connections [][]int) int {
	sort.Slice(connections, func(i, j int) bool { return connections[i][2] < connections[j][2] })

	var ans int
	uf := newUnionFind(n)
	for _, edge := range connections {
		p, q, cost := edge[0]-1, edge[1]-1, edge[2]
		if uf.connected(p, q) {
			continue
		}
		uf.union(p, q)
		ans += cost
	}

	if uf.count() == 1 {
		return ans
	}
	return -1
}

func newUnionFind(n int) *unionFind {
	ancestor := make([]int, n)
	for i := range n {
		ancestor[i] = i
	}
	return &unionFind{sz: n, ancestor: ancestor}
}

type unionFind struct {
	sz       int
	ancestor []int
}

func (o *unionFind) find(x int) int {
	if o.ancestor[x] != x {
		o.ancestor[x] = o.find(o.ancestor[x])
	}
	return o.ancestor[x]
}

func (o *unionFind) union(p, q int) {
	o.ancestor[o.find(p)] = o.find(q)
	o.sz--
}

func (o *unionFind) connected(p, q int) bool {
	return o.find(p) == o.find(q)
}

func (o *unionFind) count() int {
	return o.sz
}

func minimumCost2(n int, connections [][]int) int {
	return prim(n, connections)
}

func prim(n int, connections [][]int) int {
	graph := make([][]*pair, n)
	for _, edge := range connections {
		from, to, weight := edge[0]-1, edge[1]-1, edge[2]
		graph[from] = append(graph[from], &pair{to, weight})
	}

	var ans int
	minQ := &hp{}
	inMST := make([]bool, n)
	inMST[0] = true
	cut(0, graph, minQ, inMST)
	for minQ.Len() > 0 {
		v := heap.Pop(minQ).(*pair)
		if inMST[v.to] {
			continue
		}
		ans += v.weight
		inMST[v.to] = true
		cut(v.to, graph, minQ, inMST)
	}

	for i := range inMST {
		if !inMST[i] {
			return -1
		}
	}
	return ans
}

// 将 s 的横切边加入优先队列
func cut(s int, graph [][]*pair, minQ heap.Interface, inMST []bool) {
	for _, v := range graph[s] {
		if !inMST[v.to] {
			heap.Push(minQ, v)
		}
	}
}

type pair struct {
	to     int
	weight int
}

type hp struct{ arr []*pair }

func (h *hp) Push(v any)         { h.arr = append(h.arr, v.(*pair)) }
func (h *hp) Pop() any           { v := h.arr[len(h.arr)-1]; h.arr = h.arr[:len(h.arr)-1]; return v }
func (h *hp) Len() int           { return len(h.arr) }
func (h *hp) Less(i, j int) bool { return h.arr[i].weight < h.arr[j].weight }
func (h *hp) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func testOne(n int, in string, ans int) {
	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(minimumCost(n, matrix) == ans)
	}

	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(minimumCost2(n, matrix) == ans)
	}
}

func main() {
	testOne(3, "[[1,2,5],[1,3,6],[2,3,1]]", 6)
	testOne(4, "[[1,2,3],[3,4,4]]", -1)
}
