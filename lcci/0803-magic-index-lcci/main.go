package main

import (
	"leetcode/helper"
)

// 魔术索引
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			return i
		}
	}
	return -1
}

func findMagicIndex2(nums []int) int {
	for i := 0; i < len(nums); {
		if i == nums[i] {
			return i
		} else if nums[i] > i {
			i = nums[i]
		} else {
			i++
		}
	}
	return -1
}

func findMagicIndex3(nums []int) int {
	var f func(int, int) int

	f = func(i, j int) int {
		if i > j {
			return -1
		}
		mid := i + (j-i)/2
		if v := f(i, mid-1); v != -1 {
			return v
		}
		if nums[mid] == mid {
			return mid
		}
		return f(mid+1, j)
	}

	return f(0, len(nums)-1)
}

func testOne(nums []int, ans int) {
	helper.Assert(findMagicIndex(nums) == ans)
	helper.Assert(findMagicIndex2(nums) == ans)
	helper.Assert(findMagicIndex3(nums) == ans)
}

func main() {
	testOne([]int{0, 2, 3, 4, 5}, 0)
	testOne([]int{1, 1, 1}, 1)
}
