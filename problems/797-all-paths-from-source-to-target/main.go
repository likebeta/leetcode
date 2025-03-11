package main

import "leetcode/helper"

// 所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		traverse func([][]int, int, []int)
		ans      [][]int
	)

	traverse = func(graph [][]int, curr int, path []int) {
		path = append(path, curr)

		// 到达目标节点
		if curr == len(graph)-1 {
			newPath := append([]int{}, path...)
			ans = append(ans, newPath)
			// 由于语言特性path不需要取消
			// path = path[0 : len(path)-1]
			return
		}

		for _, p := range graph[curr] {
			traverse(graph, p, path)
		}
		// 由于语言特性path不需要取消
		// path = path[0 : len(path)-1]
	}

	traverse(graph, 0, nil)
	return ans
}

func testOne(in string, ans string) {
	graph := helper.ParseIntMatrix(in)
	paths := allPathsSourceTarget(graph)
	s := helper.DumpIntMatrix(paths)
	helper.LogF("in=%s, paths=%s, ans=%s", in, s, ans)
}

func main() {
	testOne("[[1,2],[3],[3],[]]", "[[0,1,3],[0,2,3]]")
	testOne("[[4,3,1],[3,2,4],[3],[4],[]]", "[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]")
}
