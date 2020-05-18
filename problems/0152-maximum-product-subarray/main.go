package main

import (
	"leetcode/helper"
)

// 乘积最大子数组
func maxProduct(nums []int) int {
	ans, minNum, maxNum := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minNum, maxNum = min(nums[i], nums[i]*minNum, nums[i]*maxNum), max(nums[i], nums[i]*minNum, nums[i]*maxNum)
		if maxNum > ans {
			ans = maxNum
		}
	}
	return ans
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}

func max(a, b, c int) int {
	if a < b {
		a = b
	}
	if a < c {
		return c
	}
	return a
}

func main() {
	helper.Assert(maxProduct([]int{2, 3, -2, 4}) == 6)
	helper.Assert(maxProduct([]int{-2, 0, -1}) == 0)
}
