package main

import (
	"leetcode/helper"
	"sort"
)

// 子集 II
// nums可能包含重复元素
func subsetsWithDup(nums []int) [][]int {
	var (
		ans       [][]int
		backtrack func(int)
		record    = make([]int, 0)
	)

	sort.Ints(nums)

	backtrack = func(i int) {
		one := make([]int, len(record))
		copy(one, record)
		ans = append(ans, one)

		for j := i; j < len(nums); j++ {
			// [i, len(nums)) 如果有相同元素只选择第一个， 因为选择第二个的情况被选择第一个包含了
			if j == i || nums[j] != nums[j-1] {
				record = append(record, nums[j])
				backtrack(j + 1)
				record = record[:len(record)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func main() {
	helper.Log(subsetsWithDup([]int{1, 2, 2}))
	helper.Log(subsetsWithDup([]int{0}))
}
