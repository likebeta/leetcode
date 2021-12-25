package main

import (
	"leetcode/helper"
)

// 汉明距离总和
func totalHammingDistance(nums []int) int {
	length := len(nums)
	masks := make([]int, 32, 32)
	for _, num := range nums {
		for i := 0; num > 0; i++ {
			masks[i] += num & 1
			num >>= 1
		}
	}

	var ans int
	for i := range masks {
		ans += masks[i] * (length - masks[i])
	}
	return ans
}

func main() {
	helper.Assert(totalHammingDistance([]int{4, 14, 2}) == 6)
}
