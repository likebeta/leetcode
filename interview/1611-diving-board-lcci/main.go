package main

import (
	"leetcode/helper"
)

// 跳水板
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	lengths := make([]int, k+1)
	for i := 0; i <= k; i++ {
		lengths[i] = shorter*(k-i) + longer*i
	}
	return lengths
}

func main() {
	helper.Log(divingBoard(1, 2, 3), []int{3, 4, 5, 6})
}
