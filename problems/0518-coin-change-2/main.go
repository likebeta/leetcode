package main

import (
	"leetcode/helper"
)

// 零钱兑换 II
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	N := len(coins)
	if N == 0 {
		return 0
	}
	dp := make([][]int, amount+1, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, N, N)
	}

	for j := 0; j < N; j++ {
		dp[0][j] = 1
	}

	for k := 0; k <= amount; k += coins[0] {
		dp[k][0] = 1
	}

	for i := 1; i <= amount; i++ {
		for j := 1; j < N; j++ {
			for k := 0; k <= i; k += coins[j] {
				dp[i][j] += dp[i-k][j-1]
			}
		}
	}
	return dp[amount][N-1]
}

func change2(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	N := len(coins)
	if N == 0 {
		return 0
	}
	dp := make([][]int, amount+1, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, N, N)
	}

	for j := 0; j < N; j++ {
		dp[0][j] = 1
	}

	for k := 0; k <= amount; k += coins[0] {
		dp[k][0] = 1
	}

	for i := 1; i <= amount; i++ {
		for j := 1; j < N; j++ {
			dp[i][j] = dp[i][j-1]
			if i >= coins[j] {
				dp[i][j] += dp[i-coins[j]][j]
			}
		}
	}
	return dp[amount][N-1]
}

func change3(amount int, coins []int) int {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func main() {
	var arr []int
	arr = []int{1, 2, 5}
	helper.Log(change(5, arr), change2(5, arr), change3(5, arr),  4)
	arr = []int{2}
	helper.Log(change(3, arr), change2(3, arr), change3(3, arr), 0)
	arr = []int{10}
	helper.Log(change(10, arr), change2(10, arr), change3(10, arr), 1)
}
