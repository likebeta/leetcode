package main

import (
	"leetcode/helper"
	"sort"
)

// 连接所有点的最小费用
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	var connections [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			connections = append(connections, []int{i, j, abs(xi-xj) + abs(yi-yj)})
		}
	}

	return kruskal(n, connections)
}

func kruskal(n int, connections [][]int) int {
	sort.Slice(connections, func(i, j int) bool { return connections[i][2] < connections[j][2] })

	var ans int
	uf := newUnionFind(n)
	for _, tuple := range connections {
		p, q, cost := tuple[0], tuple[1], tuple[2]
		if uf.connected(p, q) {
			continue
		}
		uf.union(p, q)
		ans += cost
	}

	if uf.count() == 1 {
		return ans
	}
	return 0
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
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

// 同1135
func prim(n int, connections [][]int) int {
	return 0
}

func testOne(in string, ans int) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(minCostConnectPoints(matrix) == ans)
}

func main() {
	testOne("[[0,0],[2,2],[3,10],[5,2],[7,0]]", 20)
	testOne("[[3,12],[-2,5],[-4,1]]", 18)
	testOne("[[0,0],[1,1],[1,0],[-1,1]]", 4)
	testOne("[[-1000000,-1000000],[1000000,1000000]]", 4000000)
	testOne("[[0,0]]", 0)
}
