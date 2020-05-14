package main

import (
	"leetcode/helper"
)

func singleNumber(nums []int) int {
	var num int
	for i := range nums {
		num ^= nums[i]
	}
	return num
}

func main() {
	helper.Assert(singleNumber([]int{2, 2, 1}) == 1)
	helper.Assert(singleNumber([]int{4, 1, 2, 1, 2}) == 4)
}
