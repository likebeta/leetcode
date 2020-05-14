package main

import (
	"leetcode/helper"
)

// 只出现一次的数字 II
/*
	once  twice  state
      0     0      0
      1     0      1
      0     1      2
*/
func singleNumber(nums []int) int {
	var once, twice int
	for i := range nums {
		// tmp := once
		// (0 0 1) || (1 0 0)
		once = (^twice) & (once ^ nums[i])
		// (1 0 1) || (0 1 0) ->   (0 0 1) || (0 1 0)
		// twice = (tmp ^ twice) & (twice ^ nums[i])
		// (1 0 1) || (0 1 0) ->   (0 0 1) || (0 1 0)
		twice = (^once) & (twice ^ nums[i])
	}
	return once
}

func main() {
	helper.Assert(singleNumber([]int{2, 2, 3, 2}) == 3)
	helper.Assert(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}) == 99)
}
