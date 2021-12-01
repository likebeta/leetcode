package main

import (
	"leetcode/helper"
)

// 二分查找
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] > target
			right = mid - 1
		}
	}
	return -1
}

func testOne(in string, target, ans int) {
	nums := helper.ParseArray(in)
	result := search(nums, target)
	helper.Assert(result == ans)
	helper.Log(in, target, "=>", ans)
}

func main() {
	testOne("[-1,0,3,5,9,12]", 9, 4)
	testOne("[-1,0,3,5,9,12]", 2, -1)
}
