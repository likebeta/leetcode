package main

import (
	"leetcode/helper"
	"math"
)

// 地下城游戏
func calculateMinimumHP(dungeon [][]int) int {
	row, col := len(dungeon), len(dungeon[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[row-1][col-1] = max(1, 1-dungeon[row-1][col-1])
	for i := col - 2; i >= 0; i-- {
		dp[row-1][i] = max(1, dp[row-1][i+1]-dungeon[row-1][i])
	}
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = max(1, dp[i+1][col-1]-dungeon[i][col-1])
	}

	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(1, dp[i][j]-dungeon[i][j])
		}
	}

	return dp[0][0]
}

func calculateMinimumHP2(dungeon [][]int) int {
	row, col := len(dungeon), len(dungeon[0])
	dp := make([]int, col+1)
	dp[col] = 1
	for i := col - 1; i >= 0; i-- {
		dp[i] = max(1, dp[i+1]-dungeon[row-1][i])
	}

	dp[col] = math.MaxInt32
	for i := row - 2; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			dp[j] = min(dp[j], dp[j+1])
			dp[j] = max(1, dp[j]-dungeon[i][j])
		}
	}

	return dp[0]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}


func testOne(matrix [][]int, ans int) {
	helper.Log(calculateMinimumHP(matrix), ans)
	helper.Log(calculateMinimumHP2(matrix), ans)
}

func main() {
	var matrix [][]int
	matrix = [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	testOne(matrix, 7)
}
