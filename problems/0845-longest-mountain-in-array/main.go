package main

import (
	"leetcode/helper"
)

// 数组中的最长山脉
func longestMountain(a []int) int {
	var (
		n    = len(a)
		left int
		ans  int
	)
	for left+2 < n {
		right := left + 1
		if a[left] < a[left+1] {
			for right+1 < n && a[right] < a[right+1] {
				right++
			}
			if right < n-1 && a[right] > a[right+1] {
				for right+1 < n && a[right] > a[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return ans
}

func main() {
	helper.Assert(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}) == 5)
	helper.Assert(longestMountain([]int{2, 2, 2}) == 0)
}
