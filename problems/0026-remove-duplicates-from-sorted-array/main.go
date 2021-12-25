package main

import (
	"leetcode/helper"
)

// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	var cnt int
	for i := 1; i < length; i++ {
		if nums[i] != nums[cnt] {
			cnt++
			nums[cnt] = nums[i]
		}
	}
	return cnt + 1
}

func main() {
	helper.Assert(removeDuplicates([]int{1, 1, 2}) == 2)
	helper.Assert(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) == 5)
}
