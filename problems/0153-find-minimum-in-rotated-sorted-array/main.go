package main

import (
	"leetcode/helper"
)

func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j && nums[i] > nums[j] {
		mid := i + (j-i)/2
		if nums[mid] >= nums[i] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return nums[i]
}

func main() {
	helper.Log(findMin([]int{2, 3, 4, 5, 1}),  1)
	helper.Log(findMin([]int{3, 4, 5, 1, 2}), 1)
	helper.Log(findMin([]int{4, 5, 6, 7, 0, 1, 2}), 0)
}
