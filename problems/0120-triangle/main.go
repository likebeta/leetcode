package main

import (
	"leetcode/helper"
	"math"
)

// 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	N := len(triangle)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < N; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < N; i++ {
		ans = min(ans, dp[N-1][i])
	}
	return ans
}

func minimumTotal2(triangle [][]int) int {
	N := len(triangle)
	dp := make([]int, N)
	dp[0] = triangle[0][0]
	for i := 1; i < N; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for i := 0; i < N; i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func testOne(matrix [][]int, ans int) {
	helper.Assert(minimumTotal(matrix) == ans)
	helper.Assert(minimumTotal2(matrix) == ans)
}

func main() {
	var matrix [][]int
	matrix = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	testOne(matrix, 11)
}
