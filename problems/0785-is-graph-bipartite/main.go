package main

import "leetcode/helper"

// 判断二分图
func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			queue := []int{i}
			color[i] = 1
			for i := 0; i < len(queue); i++ {
				node := queue[i]
				cNei := 1
				if color[node] == cNei {
					cNei = 3 - cNei
				}
				for _, neighbor := range graph[node] {
					if color[neighbor] == 0 {
						queue = append(queue, neighbor)
						color[neighbor] = cNei
					} else if color[neighbor] != cNei {
						return false
					}
				}
			}
		}
	}
	return true
}

func testOne(in string, ans bool) {
	var matrix [][]int
	helper.Load([]byte(in), &matrix)
	helper.Assert(isBipartite(matrix) == ans)
}

func main() {
	testOne("[[1,3], [0,2], [1,3], [0,2]]", true)
	testOne("[[1,2,3], [0,2], [0,1,3], [0,2]]", false)
	testOne("[[1], [0,3], [3], [1,2]]", true)
	testOne("[[3], [2,4], [1], [0,4], [1,3]]", true)
}
