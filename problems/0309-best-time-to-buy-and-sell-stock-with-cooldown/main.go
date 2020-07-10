package main

import (
	"leetcode/helper"
)

// 最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	/*
	   第n天买入的时候, 要从n-2那天的状态转移, 不再是n-1
	   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
	   dp[n].full  = max(dp[n-1].full, dp[n-2].empty - prices[n])
	*/
	N := len(prices)
	if N == 0 {
		return 0
	}
	dp := make([]Equities, N)
	dp[0].Empty = 0
	dp[0].Full = -prices[0]
	prePreEmpty := 0
	for n := 1; n < N; n++ {
		dp[n].Empty = max(dp[n-1].Empty, dp[n-1].Full+prices[n])
		// dp[n].Full = max(dp[n-1].Full, dp[n-2].Empty - prices[n])
		dp[n].Full = max(dp[n-1].Full, prePreEmpty-prices[n])
		prePreEmpty = dp[n-1].Empty
	}
	return max(dp[N-1].Full, dp[N-1].Empty)
}

type Equities struct {
	Empty int
	Full  int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	helper.Assert(maxProfit([]int{1, 2, 3, 0, 2}) == 3)
}
