package main

import (
	"leetcode/helper"
	"math"
)

// 分割数组的最大值
func splitArray(nums []int, m int) int {
	N := len(nums)
	sums := make([]int, N+1)
	for i := 0; i < N; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= i && j <= m; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sums[i]-sums[k]))
			}
		}
	}
	return dp[N][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	helper.Assert(splitArray([]int{7, 2, 5, 10, 8}, 2) == 18)
}
