package main

import (
	"leetcode/helper"
)

// 全排列
func permute(nums []int) [][]int {
	var (
		ans       [][]int
		backtrack func(int)
		record    = make([]int, len(nums))
		visited   = make([]bool, len(nums))
	)

	backtrack = func(i int) {
		if i == len(nums) {
			one := make([]int, len(record))
			copy(one, record)
			ans = append(ans, one)
			return
		}
		for j := 0; j < len(nums); j++ {
			if visited[j] {
				continue
			}
			record[i] = nums[j]
			visited[j] = true
			backtrack(i + 1)
			visited[j] = false
		}
	}
	backtrack(0)
	return ans
}

func main() {
	helper.Log(permute([]int{1, 2, 3}))
	helper.Log(permute([]int{0, 1}))
	helper.Log(permute([]int{1}))
}
