package main

import (
	"leetcode/helper"
)

// 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	helper.Log(plusOne([]int{1, 2, 3}))
	helper.Log(plusOne([]int{4, 3, 2, 1}))
	helper.Log(plusOne([]int{9}))
}
