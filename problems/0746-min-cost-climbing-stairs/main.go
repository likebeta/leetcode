package main

import (
	"leetcode/helper"
)

// 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)
	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[length]
}

func minCostClimbingStairs2(cost []int) int {
	length := len(cost)
	var ans, pre, prePre int
	for i := 2; i <= length; i++ {
		ans = min(pre+cost[i-1], prePre+cost[i-2])
		pre, prePre = ans, pre
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func testOne(arr []int, ans int) {
	helper.Assert(minCostClimbingStairs(arr) == ans)
	helper.Assert(minCostClimbingStairs2(arr) == ans)
}

func main() {
	testOne([]int{10, 15, 20}, 15)
	testOne([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6)
}
