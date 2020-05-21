package main

import (
	"leetcode/helper"
)

// 打家劫舍 II
func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	return max(doRob(nums, 0, length-1), doRob(nums, 1, length))
}

func doRob(nums []int, l, r int) int {
	var pre1, pre2, ans int
	for i := l; i < r; i++ {
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
	helper.Assert(rob([]int{2, 3, 2}) == 3)
	helper.Assert(rob([]int{1, 2, 3, 1}) == 4)
}
