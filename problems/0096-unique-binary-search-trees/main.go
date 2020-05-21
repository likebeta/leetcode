package main

import (
	"leetcode/helper"
)

// 不同的二叉搜索树
func numTrees(n int) int {
	// dp(n) = f(1) + f(2) + ... + f(i) + ... + f(n)
	// f(i) = dp(i-1)*dp(n-i)
	dp := make([]int, n+1, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {
	helper.Assert(numTrees(3) == 5)
	helper.Assert(numTrees(5) == 42)
}
