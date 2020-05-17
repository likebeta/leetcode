package main

import (
	"leetcode/helper"
)

// 移除元素
func removeElement(nums []int, val int) int {
	var index int
	for i := range nums {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeElement2(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			j--
			nums[i] = nums[j]
		} else {
			i++
		}
	}
	return j
}

func main() {
	helper.Assert(removeElement([]int{3, 2, 2, 3}, 3) == 2)
	helper.Assert(removeElement2([]int{3, 2, 2, 3}, 3) == 2)
	helper.Assert(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) == 5)
	helper.Assert(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) == 5)
}
