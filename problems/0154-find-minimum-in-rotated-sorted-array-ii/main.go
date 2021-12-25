package main

import "leetcode/helper"

// 寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] > nums[high] {
			high = mid
		} else {
			high--
		}
	}
	return nums[low]
}

func main() {
	helper.Assert(findMin([]int{3, 4, 5}) == 3)
	helper.Assert(findMin([]int{3, 4, 5, 1}) == 1)
	helper.Assert(findMin([]int{3, 4, 5, 1, 2}) == 1)
	helper.Assert(findMin([]int{2, 2, 2, 0, 1}) == 0)
}
