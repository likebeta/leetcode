package main

import (
	"leetcode/helper"
)

// 长度最小的子数组
func minSubArrayLen(s int, nums []int) int {
	l, N, sum := -1, len(nums), 0
	ans := N + 1
	for r := 0; r < N; r++ {
		if sum >= s-nums[r] {
			ans = min(ans, r-l)
			for l++; l < r; l++ {
				if sum -= nums[l]; sum < s-nums[r] {
					break
				}
				ans = min(ans, r-l)
			}
		}
		sum += nums[r]
	}

	if ans == N+1 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	helper.Log(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
	helper.Log(minSubArrayLen(3, []int{1, 1}), 0)
}
