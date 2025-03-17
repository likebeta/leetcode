package main

import (
	"leetcode/helper"
)

// 子集
func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		for j := len(result) - 1; j >= 0; j-- {
			tmp := len(result[j])
			one := make([]int, tmp+1)
			copy(one, result[j])
			one[tmp] = nums[i]
			result = append(result, one)
		}
	}
	return result
}

func subsets2(nums []int) [][]int {
	var (
		trace     []int
		result    [][]int
		backtrack func(int)
	)

	backtrack = func(start int) {
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			backtrack(i + 1)
			trace = trace[:len(trace)-1]
		}
	}

	backtrack(0)

	return result
}

func main() {
	helper.Log(subsets([]int{1, 2, 3}))
	helper.Log(subsets2([]int{1, 2, 3}))
}
