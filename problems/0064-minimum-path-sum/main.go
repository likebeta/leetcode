package main

import (
	"leetcode/helper"
)

// 最小路径和
func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([]int, col)
	dp[0] = grid[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < col; j++ {
			if dp[j] > dp[j-1] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] += grid[i][j]
			}
		}
	}
	return dp[col-1]
}

func main() {
	var matrix [][]int
	matrix = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	helper.Assert(minPathSum(matrix) == 7)
}
