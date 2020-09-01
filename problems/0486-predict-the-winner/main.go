package main

import (
	"leetcode/helper"
)

func PredictTheWinner(nums []int) bool {
	N := len(nums)
	if N < 2 || N%2 == 0 {
		return true
	}

	type State struct {
		First  int
		Second int
	}

	dp := make([][]State, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]State, N)
		dp[i][i] = State{First: nums[i], Second: 0}
	}
	// 1. 定义dp[i][j]为在[i, j]区间先手A和后手B最多拿多少
	// 2. 当A先拿nums[i], B就变成了先手，问题就变成了在[i+1, j]区间先手B和后手A最多拿多少
	// 3. 同理, 当A先拿nums[j]，问题就变成了在[i, j-1]区间先手B和后手A最多拿多少
	// 4. 取2, 3最大值就是先手A做出的最好选择
	for i := N - 2; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			dp[i][j].First = max(nums[i]+dp[i+1][j].Second, nums[j]+dp[i][j-1].Second)
			dp[i][j].Second = min(dp[i+1][j].First, dp[i][j-1].First)
		}
	}
	return dp[0][N-1].First >= dp[0][N-1].Second
}

// 预测赢家
func PredictTheWinner2(nums []int) bool {
	N := len(nums)
	if N < 2 || N%2 == 0 {
		return true
	}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = nums[i]
	}
	// 1. 定义dp[i][j]为[i, j]先手A比后手B多拿的分数
	// 2. 当A先拿nums[i], B就变成了先手, dp[i+1][j]就变成了[i+1, j]上B比A多拿的分数, nums[i] - dp[i+1][j]就是此种情况结果
	// 3. 同理, 当A先拿nums[j]， nums[j]-dp[i][j-1]就是此种情况结果
	// 4. 取2, 3最大值就是先手A做出的最好选择
	for i := N - 2; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][N-1] >= 0
}

func PredictTheWinner3(nums []int) bool {
	N := len(nums)
	if N < 2 || N%2 == 0 {
		return true
	}
	dp := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		dp[i] = nums[i]
		for j := i + 1; j < N; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[N-1] >= 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testOne(in []int, ans bool) {
	helper.Assert(PredictTheWinner(in) == ans)
	helper.Assert(PredictTheWinner2(in) == ans)
	helper.Assert(PredictTheWinner3(in) == ans)
}

func main() {
	testOne([]int{1, 5, 2}, false)
	testOne([]int{1, 5, 233, 7}, true)
}
