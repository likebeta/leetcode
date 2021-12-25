package main

import (
	"leetcode/helper"
)

// 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, col+1)
	}
	dp[0][1] = 1
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if obstacleGrid[i-1][j-1] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row][col]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, col+1)
	for j := 1; j <= col; j++ {
		if obstacleGrid[0][j-1] != 0 {
			break
		}
		dp[j] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j <= col; j++ {
			if obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[col]
}

func testOne(matrix [][]int, ans int) {
	helper.Assert(uniquePathsWithObstacles(matrix) == ans)
	helper.Assert(uniquePathsWithObstacles2(matrix) == ans)
}

func main() {
	testOne([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}, 2)

	testOne([][]int{
		{1},
		{0},
	}, 0)

	testOne([][]int{
		{0},
		{1},
	}, 0)
}
