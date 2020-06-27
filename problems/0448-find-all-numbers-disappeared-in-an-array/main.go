package main

import (
	"leetcode/helper"
)

// 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	var ans []int
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func testOne(arr []int, ans []int) {
	helper.Print(arr, " ")
	helper.Log(findDisappearedNumbers(arr), ans)
}

func main() {
	testOne([]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6})
}
