package main

import (
	"leetcode/helper"
	"slices"
)

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)

	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // 最小长度为1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

func testOne(in string, ans int) {
	nums := helper.ParseIntArray(in)
	helper.Assert(lengthOfLIS(nums) == ans)
}

func main() {
	testOne("[10,9,2,5,3,7,101,18]", 4)
	testOne("[0,1,0,3,2,3]", 4)
	testOne("[7,7,7,7,7,7,7]", 1)
}
