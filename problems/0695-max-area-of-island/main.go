package main

import "leetcode/helper"

// 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(grid, m, n, i, j))
			}
		}
	}

	return ans
}

// 返回淹掉的格子数量
func dfs(grid [][]int, m, n, i, j int) int {
	if i >= m || j >= n || i < 0 || j < 0 {
		return 0
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	return 1 + dfs(grid, m, n, i-1, j) + dfs(grid, m, n, i+1, j) + dfs(grid, m, n, i, j-1) + dfs(grid, m, n, i, j+1)
}

func testOne(in string, ans int) {
	grid := helper.ParseIntMatrix(in)
	helper.Assert(maxAreaOfIsland(grid) == ans)
}

func main() {
	testOne(`[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]`, 6)
	testOne(`[[0,0,0,0,0,0,0,0]]`, 0)
}
