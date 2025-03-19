package main

import (
	"leetcode/helper"
)

// 零钱兑换
// 状态定义：dp[i] 表示凑成金额 i 所需的最少硬币数
// 状态转移：用来凑成金额 i 的最少硬币数 = 用来凑成金额 (i-coins[j]) 的最少硬币数 + 1
// 1. 对于每个金额 i，遍历每个硬币面值 coins[j]
// 2. 计算使用该硬币后剩余的金额 pre = i - coins[j]
// 3. 如果 pre 有效且有解，那么 dp[i] = min(dp[i], dp[pre] + 1)
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

// 状态定义：dp[i][j] 的值表示使用第 j 个硬币及之后的硬币凑成金额 i 所需的最少硬币数
// 状态转移：对于金额 i 和当前硬币 j，我们有两个选择
// 1. 不使用当前硬币，直接使用后面的硬币（dp[i][j+1]）
// 2. 使用当前硬币，然后继续使用当前硬币（dp[i-coins[j]][j] + 1）
func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	N := len(coins)
	if N == 0 {
		return -1
	}
	dp := make([][]int, amount+1)
	dp[0] = make([]int, N+1) // 金额为0时，所有状态都是0
	for i := 1; i <= amount; i++ {
		dp[i] = make([]int, N+1)
		dp[i][N] = -1 // 当没有硬币可用时，无法凑成任何金额
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

func testOne(in string, amount, ans int) {
	{
		coins := helper.ParseIntArray(in)
		helper.Assert(coinChange(coins, amount) == ans)
	}

	{
		coins := helper.ParseIntArray(in)
		helper.Assert(coinChange2(coins, amount) == ans)
	}
}

func main() {
	testOne("[1,2,5]", 11, 3)
	testOne("[2]", 3, -1)
	testOne("[2]", 0, 0)
}
