package main

import (
	"leetcode/helper"
)

// 最长重复子数组
func findLength(A []int, B []int) int {
	lA, lB := len(A), len(B)
	dp := make([][]int, lA+1)
	for i := 0; i <= lA; i++ {
		dp[i] = make([]int, lB+1)
	}

	var ans int
	for i := 1; i <= lA; i++ {
		for j := 1; j <= lB; j++ {
			if A[i-1] == B[j-1] {
				if dp[i][j] = dp[i-1][j-1] + 1; dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
}

func findLength2(A []int, B []int) int {
	lA, lB := len(A), len(B)
	if lA > lB {
		return findLength2(B, A)
	}
	dp := make([]int, lA+1)
	for j := 1; j <= lB; j++ {
		var pre int
		for i := 1; i <= lA; i++ {
			pre, dp[i] = dp[i], pre
			if A[i-1] == B[j-1] {
				if dp[i] = dp[i] + 1; dp[i] > dp[0] {
					dp[0] = dp[i]
				}
			} else {
				dp[i] = 0
			}
		}
	}
	return dp[0]
}

func testOne(a, b []int, ans int) {
	helper.Log(findLength(a, b), ans)
	helper.Log(findLength2(a, b), ans)
}

func main() {
	testOne([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3)
}
