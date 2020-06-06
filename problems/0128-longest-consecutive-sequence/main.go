package main

import (
	"leetcode/helper"
)

// 最长连续序列
func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for i := range nums {
		set[nums[i]] = true
	}

	var c, ans int
	for num := range set {
		if !set[num-1] {
			for c = 0; set[num]; num++ {
				c++
			}
			if c > ans {
				ans = c
			}
		}
	}
	return ans
}

func main() {
	helper.Assert(longestConsecutive([]int{}) == 0)
	helper.Assert(longestConsecutive([]int{100}) == 1)
	helper.Assert(longestConsecutive([]int{100, 1}) == 1)
	helper.Assert(longestConsecutive([]int{100, 4, 200, 1, 3, 2}) == 4)
}
