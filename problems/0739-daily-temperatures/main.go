package main

import (
	"leetcode/helper"
)

// 每日温度
func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := make([]int, len(T)+1)
	stack[0] = -1
	var k int
	for i := range T {
		for top := stack[k]; top != -1 && T[top] < T[i]; top = stack[k] {
			ans[top], k = i-top, k-1
		}
		k++
		stack[k] = i
	}
	return ans
}

func main() {
	helper.Log(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})
}
