package main

import (
	"leetcode/helper"
)

// 以图判树
func validTree(n int, edges [][]int) bool {
	uf := newUnionFind(n)
	for _, pair := range edges {
		p, q := pair[0], pair[1]
		if uf.connected(p, q) {
			return false
		}
		uf.union(p, q)
	}

	return uf.count() == 1
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

func testOne(n int, in string, ans bool) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(validTree(n, matrix) == ans)
}

func main() {
	testOne(5, "[[0,1],[0,2],[0,3],[1,4]]", true)
	testOne(5, "[[0,1],[1,2],[2,3],[1,3],[1,4]]", false)
}
