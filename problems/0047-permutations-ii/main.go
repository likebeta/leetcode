package main

import (
	"leetcode/helper"
	"sort"
)

// 全排列 II
func permuteUnique(nums []int) [][]int {
	var (
		ans       [][]int
		backtrack func(int)
		record    = make([]int, len(nums))
		visited   = make([]bool, len(nums))
	)

	sort.Ints(nums)

	backtrack = func(i int) {
		if i == len(nums) {
			one := make([]int, len(record))
			copy(one, record)
			ans = append(ans, one)
			return
		}
		for j := 0; j < len(nums); j++ {
			if visited[j] || (j > 0 && nums[j] == nums[j-1] && !visited[j-1]) {
				continue
			}
			record[i] = nums[j]
			visited[j] = true
			backtrack(i+1)
			visited[j] = false
		}
	}
	backtrack(0)
	return ans
}

func main() {
	helper.Log(permuteUnique([]int{1, 1, 2}))
}
