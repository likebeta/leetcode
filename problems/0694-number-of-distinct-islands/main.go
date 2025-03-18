package main

import "leetcode/helper"

// 不同岛屿的数量
func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	paths := make(map[string]bool)
	path := []byte{}
	var dfs func(int, int, byte)

	dfs = func(i, j int, dir byte) {
		if i >= m || j >= n || i < 0 || j < 0 {
			return
		}

		if grid[i][j] == 1 {
			path = append(path, dir)
			grid[i][j] = 0
			dfs(i-1, j, 1)
			dfs(i+1, j, 2)
			dfs(i, j-1, 3)
			dfs(i, j+1, 4)
		}
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				dfs(i, j, 0)
				paths[string(path)] = true
				path = []byte{}
			}
		}
	}

	return len(paths)
}

func testOne(in string, ans int) {
	grid := helper.ParseIntMatrix(in)
	helper.Assert(numDistinctIslands(grid) == ans)
}

func main() {
	testOne(`[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]`, 1)
	testOne(`[[1,1,0,1,1],[1,0,0,0,0],[0,0,0,0,1],[1,1,0,1,1]]`, 3)
}
