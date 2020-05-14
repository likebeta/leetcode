package main

import (
	"leetcode/helper"
)

// 缺失数字
func missingNumber(nums []int) int {
	xor := len(nums)
	for i := range nums {
		xor ^= i
		xor ^= nums[i]
	}
	return xor
}

func main() {
	helper.Assert(missingNumber([]int{3, 0, 1}) == 2)
	helper.Assert(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)
}
