package main

import (
	"leetcode/helper"
)

// 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	helper.Assert(firstMissingPositive([]int{1, 2, 0}) == 3)
	helper.Assert(firstMissingPositive([]int{3, 4, -1, 1}) == 2)
	helper.Assert(firstMissingPositive([]int{7, 8, 9, 11, 12}) == 1)
}
