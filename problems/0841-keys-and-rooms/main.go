package main

import (
	"leetcode/helper"
)

// 钥匙和房间
func canVisitAllRooms(rooms [][]int) bool {
	var (
		visitedCount int
		visited      []bool
		N            = len(rooms)
		dfs          func(i int)
	)
	visited = make([]bool, N)
	dfs = func(i int) {
		visited[i] = true
		visitedCount++
		for _, k := range rooms[i] {
			if !visited[k] {
				dfs(k)
			}
		}
	}
	dfs(0)
	return visitedCount == N
}

func canVisitAllRooms2(rooms [][]int) bool {
	var (
		N            = len(rooms)
		visited      []bool
		queue        = []int{0}
		visitedCount = 1
	)
	visited = make([]bool, N)
	visited[0] = true
	for i := 0; i < len(queue); i++ {
		for _, it := range rooms[queue[i]] {
			if !visited[it] {
				visitedCount++
				visited[it] = true
				queue = append(queue, it)
			}
		}
	}
	return visitedCount == N
}

func testOne(s [][]int, ans bool) {
	helper.Assert(canVisitAllRooms(s) == ans)
	helper.Assert(canVisitAllRooms2(s) == ans)
}

func main() {
	testOne([][]int{{1}, {2}, {3}, {}}, true)
	testOne([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false)
}
