package main

import (
	"leetcode/helper"
)

// 不同的二叉搜索树
func numTrees(n int) int {
	// dp(n) = f(1) + f(2) + ... + f(i) + ... + f(n)
	// f(i) = dp(i-1)*dp(n-i)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func numTrees2(n int) int {
	memo := make(map[int]int, n)
	return count(memo, 1, n)
}

func count(memo map[int]int, i, j int) int {
	if i >= j {
		return 1
	}
	if v, ok := memo[j-i]; ok {
		return v
	}
	var ans int
	for k := i; k <= j; k++ {
		left := count(memo, i, k-1)
		right := count(memo, k+1, j)
		ans += left * right
	}
	memo[j-i] = ans
	return ans
}

func main() {
	helper.Assert(numTrees(3) == 5)
	helper.Assert(numTrees(5) == 42)

	helper.Assert(numTrees2(3) == 5)
	helper.Assert(numTrees2(5) == 42)
}
