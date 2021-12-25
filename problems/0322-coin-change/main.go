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
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for j := range coins {
			pre := i - coins[j]
			if pre >= 0 && dp[pre] != -1 && (dp[i] == -1 || dp[i] > dp[pre]+1) {
				dp[i] = dp[pre] + 1
			}
		}
	}
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	N := len(coins)
	if N == 0 {
		return -1
	}
	dp := make([][]int, amount+1, amount+1)
	dp[0] = make([]int, N+1, N+1)
	for i := 1; i <= amount; i++ {
		dp[i] = make([]int, N+1, N+1)
		dp[i][N] = -1
	}

	for i := 1; i <= amount; i++ {
		for j := N - 1; j >= 0; j-- {
			// dp[i][j] = -1
			/*   dp[i][j] = min{dp[i-k*coins[j]][j+1]+k}  0 <= k <= i/coins[j] */
			// var coin int
			// for k := 0; coin <= i; k++ {
			// 	if dp[i-coin][j+1] != -1 && (dp[i][j] == -1 || dp[i][j] > dp[i-coin][j+1]+k) {
			// 		dp[i][j] = dp[i-coin][j+1] + k
			// 	}
			// 	coin += coins[j]
			// }
			/*   dp[i-coins[j]][j] = min{dp[i-k*coins[j]][j+1]+k-1}  1 <= k <= i/coins[j] */
			// dp[i[j] = min{dp[i-coins[j]][j]+1, dp[i][j+1]}
			dp[i][j] = dp[i][j+1]
			if i >= coins[j] && dp[i-coins[j]][j] != -1 && (dp[i][j] == -1 || dp[i][j] > dp[i-coins[j]][j]+1) {
				dp[i][j] = dp[i-coins[j]][j] + 1
			}
			// 可以继续优化为1维
		}
	}
	return dp[amount][0]
}

func main() {
	var arr []int
	arr = []int{1, 2, 5}
	helper.Log(coinChange(arr, 11), coinChange2(arr, 11), 3)
	arr = []int{2}
	helper.Log(coinChange(arr, 3), coinChange2(arr, 3), -1)
}
