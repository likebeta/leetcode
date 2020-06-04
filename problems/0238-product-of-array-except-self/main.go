package main

import (
	"leetcode/helper"
)

// 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	ans := make([]int, length, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	r := nums[length-1]
	for i := length - 2; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}

func main() {
	helper.Log(productExceptSelf([]int{2}), []int{1})
	helper.Log(productExceptSelf([]int{3, 5}), []int{5, 3})
	helper.Log(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
}
