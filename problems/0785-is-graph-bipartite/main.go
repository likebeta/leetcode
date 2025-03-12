package main

import (
	"leetcode/helper"
)

// 判断二分图
func isBipartite(graph [][]int) bool {
	var (
		ok     = true
		n      = len(graph)
		colors = make([]int, n) // 0表示未染色，1和2表示不同颜色
	)

	for vertex := 0; vertex < n && ok; vertex++ {
		if colors[vertex] == 0 {
			ok = bfs(vertex, graph, colors)
		}
	}

	return ok
}

func bfs(vertex int, graph [][]int, colors []int) bool {
	queue := []int{vertex}
	colors[vertex] = 1
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		cNei := 3 - colors[node] // 1 or 2
		for _, neighbor := range graph[node] {
			if colors[neighbor] == 0 {
				colors[neighbor] = cNei
				queue = append(queue, neighbor)
			} else if colors[neighbor] != cNei {
				return false
			}
		}
	}

	return true
}

func isBipartite2(graph [][]int) bool {
	var (
		ok     = true
		n      = len(graph)
		colors = make([]int, n) // 0表示未染色，1和2表示不同颜色
	)

	for vertex := 0; vertex < n && ok; vertex++ {
		if colors[vertex] == 0 {
			colors[vertex] = 1
			ok = dfs(vertex, graph, colors)
		}
	}

	return ok
}

func dfs(vertex int, graph [][]int, colors []int) bool {
	cNei := 3 - colors[vertex] // 1 or 2
	for _, neighbor := range graph[vertex] {
		if colors[neighbor] == 0 {
			colors[neighbor] = cNei
			if !dfs(neighbor, graph, colors) {
				return false
			}
		} else if colors[neighbor] != cNei {
			return false
		}
	}
	return true
}

func testOne(in string, ans bool) {
	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(isBipartite(matrix) == ans)
	}

	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(isBipartite2(matrix) == ans)
	}
}

func main() {
	testOne("[[1,3], [0,2], [1,3], [0,2]]", true)
	testOne("[[1,2,3], [0,2], [0,1,3], [0,2]]", false)
	testOne("[[1], [0,3], [3], [1,2]]", true)
	testOne("[[3], [2,4], [1], [0,4], [1,3]]", true)
}
