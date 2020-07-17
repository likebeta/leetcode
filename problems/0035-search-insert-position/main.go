package main

import (
	"leetcode/helper"
)

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	N := len(nums)
	left, right := 0, N-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	helper.Assert(searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	helper.Assert(searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	helper.Assert(searchInsert([]int{1, 3, 5, 6}, 4) == 2)
	helper.Assert(searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	helper.Assert(searchInsert([]int{1, 3, 5, 6}, 0) == 0)
}
