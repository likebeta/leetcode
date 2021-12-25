package main

import (
	"leetcode/helper"
)

// 两数之和
func twoSum(nums []int, target int) []int {
	mm := make(map[int]int)
	for i, v := range nums {
		if n, ok := mm[target-v]; ok {
			return []int{n, i}
		}
		mm[v] = i
	}
	return nil
}

func main() {
	helper.Log(twoSum([]int{2, 7, 11, 15}, 9))
}
