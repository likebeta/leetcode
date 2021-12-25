package main

import (
	"leetcode/helper"
)

// 子集
func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	result := [][]int{{}}
	for i := 0; i < length; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			tmp := len(result[j])
			one := make([]int, tmp+1, tmp+1)
			one[tmp] = nums[i]
			copy(one, result[j])
			result = append(result, one)
		}
	}
	return result
}

func main() {
	helper.Log(subsets([]int{1, 2, 3}))
}
