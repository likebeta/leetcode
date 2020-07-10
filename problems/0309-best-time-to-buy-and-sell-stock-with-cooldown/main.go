package main

import (
	"leetcode/helper"
)

// 最佳买卖股票时机含冷冻期
/*
   第n天买入的时候, 要从n-2那天的状态转移, 不再是n-1
   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
   dp[n].full  = max(dp[n-1].full, dp[n-2].empty - prices[n])
*/
func maxProfit(prices []int) int {
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
	return dp[N-1].Empty
}

func maxProfit2(prices []int) int {
	N := len(prices)
	if N == 0 {
		return 0
	}
	prePreEmpty := 0
	pre := Equities{Empty: 0, Full: -prices[0]}
	for n := 1; n < N; n++ {
		tmp := pre
		pre.Empty = max(tmp.Empty, tmp.Full+prices[n])
		pre.Full = max(tmp.Full, prePreEmpty-prices[n])
		prePreEmpty = tmp.Empty
	}
	return pre.Empty
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

func testOne(arr []int, ans int) {
	helper.Assert(maxProfit(arr) == ans)
	helper.Assert(maxProfit2(arr) == ans)
}

func main() {
	testOne([]int{1, 2, 3, 0, 2}, 3)
}
