package main

import (
	"leetcode/helper"
)

// 课程表
func findOrder(numCourses int, prerequisites [][]int) bool { // 深度优先
	var (
		hasCycle bool
		dfs      func(int) bool
		edges    = make([][]int, numCourses) // 邻接表，表示图的结构
		visited  = make([]int, numCourses)   // 记录节点访问状态：0: 未访问 1: 正在访问（当前DFS路径中）2: 已完成访问
	)

	for i := range prerequisites {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
	}

	dfs = func(c int) bool {
		visited[c] = 1
		for i := range edges[c] {
			if visited[edges[c][i]] == 0 {
				if dfs(edges[c][i]) {
					return true
				}
			} else if visited[edges[c][i]] == 1 {
				return true
			}
		}
		visited[c] = 2
		return false
	}
	for i := 0; i < numCourses && !hasCycle; i++ {
		if visited[i] == 0 {
			hasCycle = dfs(i)
		}
	}
	return !hasCycle
}

func findOrder2(numCourses int, prerequisites [][]int) bool { // 广度优先
	edges := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for i := range prerequisites {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++
	}
	var queue []int
	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := len(queue)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, nc := range edges[c] {
			inDegree[nc]--
			if inDegree[nc] == 0 {
				queue = append(queue, nc)
				count++
			}
		}
	}
	return count == numCourses
}

func testOne(num int, in string, ans bool) {
	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(findOrder(num, matrix) == ans)
	}

	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(findOrder2(num, matrix) == ans)
	}
}

func main() {
	testOne(2, "[[1,0]]", true)
	testOne(4, "[[1,0],[2,0],[3,1],[3,2]]", true)
	testOne(2, "[[1,0],[0, 1]]", false)
}
