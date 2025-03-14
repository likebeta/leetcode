package main

import (
	"container/heap"
	"leetcode/helper"
)

// 最小体力消耗路径
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	costs := dijkstra(heights, m, n)
	return costs[m-1][n-1]
}

func dijkstra(heights [][]int, m, n int) [][]int {
	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = -1
		}
	}

	dirs := [5]int{-1, 0, 1, 0, -1}

	minQ := &hp{}
	heap.Push(minQ, &pair{0, 0, 0})
	costs[0][0] = 0
	for minQ.Len() > 0 {
		from := heap.Pop(minQ).(*pair)

		if from.x == m-1 && from.y == n-1 {
			break
		}

		for k := 0; k < 4; k++ {
			x, y := from.x+dirs[k], from.y+dirs[k+1]
			if x >= 0 && x < m && y >= 0 && y < n {
				// 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的
				d := max(costs[from.x][from.y], abs(heights[from.x][from.y]-heights[x][y]))
				if costs[x][y] == -1 || d < costs[x][y] {
					costs[x][y] = d
					heap.Push(minQ, &pair{x, y, d})
				}
			}
		}
	}

	return costs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type pair struct {
	x, y, w int
}

type hp struct{ arr []*pair }

func (h *hp) Push(v any)         { h.arr = append(h.arr, v.(*pair)) }
func (h *hp) Pop() any           { v := h.arr[len(h.arr)-1]; h.arr = h.arr[:len(h.arr)-1]; return v }
func (h *hp) Len() int           { return len(h.arr) }
func (h *hp) Less(i, j int) bool { return h.arr[i].w < h.arr[j].w }
func (h *hp) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func testOne(in string, ans int) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(minimumEffortPath(matrix) == ans)
}

func main() {
	testOne("[[1,2,2],[3,8,2],[5,3,5]]", 2)
	testOne("[[1,2,3],[3,8,4],[5,3,5]]", 1)
	testOne("[[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]", 0)
}
