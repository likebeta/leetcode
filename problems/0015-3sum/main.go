package main

import (
	"leetcode/helper"
	"sort"
)

// 三数之和
func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := length - 1
		for l := i + 1; l < r; l++ {
			if l > i+1 && nums[l] == nums[l-1] {
				continue
			}

			for l < r && nums[l]+nums[r] > -nums[i] {
				r--
			}

			if l != r && nums[l]+nums[r] == -nums[i] {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
			}
		}
	}
	return ans
}

func main() {
	var nums []int
	var ans [][]int
	nums = []int{-1, 0, 1, 2, -1, -4}
	ans = [][]int{
		{-1, 0, 1},
		{-1, -1, 2},
	}
	helper.Log(threeSum(nums), ans)
}
