package main

import "leetcode/helper"

// 统计封闭岛屿的数目
func closedIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	// 淹掉第一列和最后一列
	for i := 0; i < m; i++ {
		dfs(grid, m, n, i, 0)
		dfs(grid, m, n, i, n-1)
	}
	// 淹掉第一行和最后一行
	for j := 1; j < n-1; j++ {
		dfs(grid, m, n, 0, j)
		dfs(grid, m, n, m-1, j)
	}

	// 淹掉封闭岛屿
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				ans++
				dfs(grid, m, n, i, j)
			}
		}
	}

	return ans
}

func dfs(grid [][]int, m, n, i, j int) {
	if i >= m || j >= n || i < 0 || j < 0 {
		return
	}

	if grid[i][j] == 0 {
		grid[i][j] = 1
		dfs(grid, m, n, i-1, j)
		dfs(grid, m, n, i+1, j)
		dfs(grid, m, n, i, j-1)
		dfs(grid, m, n, i, j+1)
	}
}

func testOne(in string, ans int) {
	grid := helper.ParseIntMatrix(in)
	helper.Assert(closedIsland(grid) == ans)
}

func main() {
	testOne(`[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]`, 2)
	testOne(`[[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]`, 1)
	testOne(`[[1,1,1,1,1,1,1],[1,0,0,0,0,0,1],[1,0,1,1,1,0,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,0,0,0,0,0,1],[1,1,1,1,1,1,1]]`, 2)
}
