package main

import (
	"leetcode/helper"
)

// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left := leftBound(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := rightBound(nums, target)
	return []int{left, right}
}

// 二分搜索最左出现
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] > target
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 二分搜索最右出现
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] > target
			right = mid - 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func testOne(in string, target int, ans string) {
	arr := helper.ParseArray(in)
	result := searchRange(arr, target)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("[5,7,7,8,8,10]", 8, "[3,4]")
	testOne("[5,7,7,8,8,10]", 6, "[-1,-1]")
	testOne("[]", 0, "[-1,-1]")
}
