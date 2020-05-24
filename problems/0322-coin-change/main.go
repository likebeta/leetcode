package main

import (
	"leetcode/helper"
)

// 零钱兑换
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for k := 1; k <= amount; k++ {
		dp[k] = -1
		for j := range coins {
			pre := k - coins[j]
			if pre >= 0 && dp[pre] != -1 && (dp[k] == -1 || dp[k] > dp[pre]+1) {
				dp[k] = dp[pre] + 1
			}
		}
	}
	return dp[amount]
}

func main() {
	helper.Assert(coinChange([]int{1, 2, 5}, 11) == 3)
	helper.Assert(coinChange([]int{2}, 3) == -1)
}
