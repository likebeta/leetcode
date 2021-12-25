package main

import (
	"leetcode/helper"
)

// 计数二进制子串
func countBinarySubstrings(s string) int {
	var ans int
	if sLen := len(s); sLen > 1 {
		stat := 1
		count := [2]int{0, 0}
		count[s[0]-'0'] = 1
		for i := 1; i < sLen; i++ {
			if s[i] != s[i-1] {
				stat++
			}
			if stat > 2 {
				ans += min(count[0], count[1])
				count[s[i] - '0'] = 0
				stat = 2
			}
			count[s[i]-'0']++
		}
		ans += min(count[0], count[1])
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
	helper.Assert(countBinarySubstrings("00110011") == 6)
	helper.Assert(countBinarySubstrings("10101") == 4)
}
