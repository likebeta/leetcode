package main

import (
	"leetcode/helper"
)

// 交错字符串
func isInterleave(s1 string, s2 string, s3 string) bool {
	N, M, T := len(s1), len(s2), len(s3)
	if (N + M) != T {
		return false
	}
	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, M+1)
	}
	dp[0][0] = true
	for i := 0; i <= N; i++ {
		for j := 0; j <= M; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[N][M]
}

func main() {
	helper.Assert(isInterleave("aabcc", "dbbca", "aadbbcbcac") == true)
	helper.Assert(isInterleave("aabcc", "dbbca", "aadbbbaccc") == false)
}
