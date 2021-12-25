package main

import (
	"leetcode/helper"
)

// 打家劫舍
func rob(nums []int) int {
	length := len(nums)
	var dp1, dp2, ans int
	for i := length - 1; i >= 0; i-- {
		ans = max(nums[i]+dp2, dp1)
		dp1, dp2 = ans, dp1
	}
	return ans
}

func rob2(nums []int) int {
	length := len(nums)
	var pre1, pre2, ans int
	for i := 0; i < length; i++ {
		ans = max(nums[i]+pre2, pre1)
		pre1, pre2 = ans, pre1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var arr []int
	arr = []int{1, 2, 3, 1}
	helper.Log(rob(arr), rob2(arr), 4)
	arr = []int{2, 7, 9, 3, 1}
	helper.Log(rob(arr), rob2(arr), 12)
}
