package main

import (
	"leetcode/helper"
)

// 移动零
func moveZeroes(nums []int) {
	var k int
	for i := range nums {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for k < len(nums) {
		nums[k] = 0
		k++
	}
}

func testOne(in string) {
	arr := helper.ParseArray(in)
	moveZeroes(arr)
	helper.Log(in, helper.DumpArray(arr))
}

func main() {
	testOne("[0, 1, 0, 3, 12]")
}
