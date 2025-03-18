package main

import "leetcode/helper"

// 岛屿数量
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, m, n, i, j)
			}
		}
	}

	return ans
}

func dfs(grid [][]byte, m, n, i, j int) {
	if i >= m || j >= n || i < 0 || j < 0 {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, m, n, i-1, j)
		dfs(grid, m, n, i+1, j)
		dfs(grid, m, n, i, j-1)
		dfs(grid, m, n, i, j+1)
	}
}

func testOne(in string, ans int) {
	matrix := helper.ParseMatrix[string](in)
	grid := make([][]byte, len(matrix))
	for i := range grid {
		grid[i] = make([]byte, len(matrix[0]))
		for j := range grid[0] {
			grid[i][j] = matrix[i][j][0]
		}
	}
	helper.Assert(numIslands(grid) == ans)
}

func main() {
	testOne(`[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`, 1)
	testOne(`[["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`, 3)
}
