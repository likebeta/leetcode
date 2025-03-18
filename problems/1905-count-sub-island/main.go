package main

import "leetcode/helper"

// 统计子岛屿
func countSubIslands(grid1 [][]int, grid2 [][]int) (ans int) {
	m, n := len(grid2), len(grid2[0])

	// grid2中是陆地，grid1中不是， 该岛屿一定不是子岛屿
	for i := range m {
		for j := range n {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, m, n, i, j)
			}
		}
	}

	// 只算岛屿数量
	for i := range m {
		for j := range n {
			if grid2[i][j] == 1 {
				ans++
				dfs(grid2, m, n, i, j)
			}
		}
	}

	return ans
}

func dfs(grid [][]int, m, n, i, j int) {
	if i >= m || j >= n || i < 0 || j < 0 {
		return
	}

	if grid[i][j] == 1 {
		grid[i][j] = 0
		dfs(grid, m, n, i-1, j)
		dfs(grid, m, n, i+1, j)
		dfs(grid, m, n, i, j-1)
		dfs(grid, m, n, i, j+1)
	}
}

func testOne(in1, in2 string, ans int) {
	grid1 := helper.ParseIntMatrix(in1)
	grid2 := helper.ParseIntMatrix(in2)
	helper.Assert(countSubIslands(grid1, grid2) == ans)
}

func main() {
	testOne(`[[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]`, `[[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]`, 3)
	testOne(`[[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]]`, `[[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]`, 2)
}
