package main

import (
	"leetcode/helper"
)

// 只出现一次的数字 III
func singleNumber(nums []int) []int {
	ans := make([]int, 2, 2)
	for i := range nums {
		ans[0] ^= nums[i]
	}

	flagBit := ans[0] & (-ans[0])
	for i := range nums {
		if nums[i]&flagBit == flagBit {
			ans[1] ^= nums[i]
		}
	}
	ans[0] ^= ans[1]
	return ans
}

func main() {
	helper.Log(singleNumber([]int{1, 2, 1, 3, 2, 5})) // [3, 5]
}
