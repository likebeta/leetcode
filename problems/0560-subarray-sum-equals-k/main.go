package main

import (
	"leetcode/helper"
)

// 和为K的子数组
func subarraySum(nums []int, k int) int {
	sumMap := map[int]int{0: 1}
	var count, sum int
	for i := range nums {
		sum += nums[i]
		count += sumMap[sum-k]
		sumMap[sum]++
	}
	return count
}

func main() {
	helper.Log(subarraySum([]int{1, 1, 1}, 2))
}
