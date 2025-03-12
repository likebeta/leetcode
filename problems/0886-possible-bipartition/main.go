package main

import "leetcode/helper"

// 可能的二分法
func possibleBipartition(n int, dislikes [][]int) bool {
	var (
		ok     = true
		colors = make([]int, n)          // 0表示未染色，1和2表示不同颜色
		graph  = buildGraph(n, dislikes) // 邻接表
	)

	for vertex := 0; vertex < n && ok; vertex++ {
		if colors[vertex] == 0 {
			colors[vertex] = 1
			ok = dfs(vertex, graph, colors)
		}
	}

	return ok
}

func buildGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n)
	for _, pair := range dislikes {
		i, j := pair[0]-1, pair[1]-1
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	return graph
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

func testOne(n int, in string, ans bool) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(possibleBipartition(n, matrix) == ans)
}

func main() {
	testOne(4, "[[1,2],[1,3],[2,4]]", true)
	testOne(3, "[[1,2],[1,3],[2,3]]", false)
	testOne(5, "[[1,2],[2,3],[3,4],[4,5],[1,5]]", false)
}
