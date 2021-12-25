package main

import (
	"leetcode/helper"
)

// 最佳观光组合
func maxScoreSightseeingPair(A []int) int {
	var ans, tmp int
	length := len(A)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if tmp = A[i] + A[j] + i - j; tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}

func maxScoreSightseeingPair2(A []int) int {
	// A[i] + A[j] + i - j = (A[i] + i) + (A[j] - j)
	var ans, tmp int
	length, preMax := len(A), A[0]
	for j := 1; j < length; j++ {
		if tmp = A[j] - j + preMax; tmp > ans {
			ans = tmp
		}
		if tmp = A[j] + j; tmp > preMax {
			preMax = tmp
		}
	}
	return ans
}

func main() {
	var arr []int
	arr = []int{8, 1, 5, 2, 6}
	helper.Log(maxScoreSightseeingPair(arr), maxScoreSightseeingPair2(arr), 11)
}
