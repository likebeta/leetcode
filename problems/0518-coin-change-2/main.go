package main

import (
	"leetcode/helper"
)

// 零钱兑换 II
// 给定不同面额的硬币和一个总金额，计算可以凑成总金额的硬币组合数
// 每一种面额的硬币有无限个
func change(amount int, coins []int) int {
	// 如果总金额为0，只有一种组合方式，即不选任何硬币
	if amount == 0 {
		return 1
	}
	// 获取硬币种类
	N := len(coins)
	// 如果没有硬币，则无法组成任何金额
	if N == 0 {
		return 0
	}

	// 创建二维DP数组，dp[i][j]表示使用coins[0...j]硬币组成金额i的方法数
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, N)
	}

	// 组成金额0只有一种方法（不选任何硬币）
	for j := 0; j < N; j++ {
		dp[0][j] = 1
	}

	// 使用第一枚硬币只能组成金额0, coins[0], 2*coins[0], ...
	for k := 0; k <= amount; k += coins[0] {
		dp[k][0] = 1
	}

	// 填充DP数组
	for i := 1; i <= amount; i++ {
		for j := 1; j < N; j++ {
			// 对于当前硬币coins[j]，考虑使用0个、1个、2个...直到不能再使用
			for k := 0; k <= i; k += coins[j] {
				dp[i][j] += dp[i-k][j-1]
			}
		}
	}
	return dp[amount][N-1]
}

// 优化版本1：改进状态转移方程
func change2(amount int, coins []int) int {
	// 基本情况处理与第一个函数相同
	if amount == 0 {
		return 1
	}
	N := len(coins)
	if N == 0 {
		return 0
	}

	// 创建并初始化DP数组
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, N)
	}

	// 初始化：组成金额0只有一种方法
	for j := 0; j < N; j++ {
		dp[0][j] = 1
	}

	// 使用第一枚硬币初始化
	for k := 0; k <= amount; k += coins[0] {
		dp[k][0] = 1
	}

	// 优化的状态转移方程
	for i := 1; i <= amount; i++ {
		for j := 1; j < N; j++ {
			if i < coins[j] {
				// 不使用当前硬币的方法数
				dp[i][j] = dp[i][j-1]
			} else {
				// 如果当前金额大于等于当前硬币面值，可以使用当前硬币或者不使用当前硬币
				dp[i][j] = dp[i-coins[j]][j] + dp[i][j-1]
			}
		}
	}
	return dp[amount][N-1]
}

// 优化版本2：空间复杂度优化，使用一维数组
func change3(amount int, coins []int) int {
	// 创建一维DP数组，dp[i]表示组成金额i的方法数
	dp := make([]int, amount+1)
	// 组成金额0只有一种方法（不选任何硬币）
	dp[0] = 1

	// 遍历每种硬币
	for _, coin := range coins {
		// 对于每种硬币，更新能组成的金额
		// 注意这里是从coin开始，因为小于coin的金额不可能使用当前硬币
		for i := coin; i <= amount; i++ {
			// 状态转移：当前金额的方法数 += 当前金额减去当前硬币值的方法数
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func testOne(in string, amount int, ans int) {
	{
		coins := helper.ParseIntArray(in)
		helper.Assert(change(amount, coins) == ans)
	}

	{
		coins := helper.ParseIntArray(in)
		helper.Assert(change2(amount, coins) == ans)
	}

	{
		coins := helper.ParseIntArray(in)
		helper.Assert(change3(amount, coins) == ans)
	}
}

func main() {
	testOne("[1,2,5]", 5, 4)
	testOne("[2]", 3, 0)
	testOne("[10]", 10, 1)
}
