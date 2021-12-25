package main

import (
	"leetcode/helper"
)

// 课程表 II
func findOrder(numCourses int, prerequisites [][]int) []int { // 深度优先
	var (
		ans      []int
		hasCycle bool
		dfs      func(int) bool
	)

	edges := make([][]int, numCourses, numCourses)
	visited := make([]int, numCourses, numCourses)
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
		ans = append(ans, c)
		return false
	}
	for i := 0; i < numCourses && !hasCycle; i++ {
		if visited[i] == 0 {
			hasCycle = dfs(i)
		}
	}

	if len(ans) != numCourses {
		return nil
	}

	for i := 0; i < numCourses; i++ {
		numCourses--
		ans[i], ans[numCourses] = ans[numCourses], ans[i]
	}
	return ans
}

func findOrder2(numCourses int, prerequisites [][]int) []int { // 广度优先
	var ans []int
	edges := make([][]int, numCourses, numCourses)
	inDegree := make([]int, numCourses, numCourses)
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
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		ans = append(ans, c)
		for _, nc := range edges[c] {
			inDegree[nc]--
			if inDegree[nc] == 0 {
				queue = append(queue, nc)
			}
		}
	}
	if len(ans) != numCourses {
		return nil
	}
	return ans
}

func main() {
	var pq [][]int
	pq = [][]int{{1, 0}}
	helper.Log(findOrder(2, pq), findOrder2(2, pq))
	pq = [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	helper.Log(findOrder(4, pq), findOrder2(4, pq))
}
