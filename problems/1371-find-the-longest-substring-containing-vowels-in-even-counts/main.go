package main

import (
	"leetcode/helper"
)

// 每个元音包含偶数次的最长子字符串
func findTheLongestSubstring(s string) int {
	cache := make([]int, 1<<5)
	for i := range cache {
		cache[i] = -1
	}
	var ans, status int
	for i := range s {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if status == 0 {
			ans = i + 1
		} else if cache[status] >= 0 {
			if i-cache[status] > ans {
				ans = i - cache[status]
			}
		} else {
			cache[status] = i
		}
	}
	return ans
}

func main() {
	helper.Log(findTheLongestSubstring("eleetminicoworoep"), 13)
	helper.Log(findTheLongestSubstring("leetcodeisgreat"), 5)
	helper.Log(findTheLongestSubstring("bcbcbc"), 6)
}
