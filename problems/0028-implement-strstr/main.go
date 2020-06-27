package main

import (
	"leetcode/helper"
)

// 实现 strStr()
func strStr(haystack string, needle string) int {
	L, N := len(needle), len(haystack)
	for i := 0; i < N-L+1; i++ {
		if haystack[i:i+L] == needle {
			return i
		}
	}
	return -1
}

func main() {
	helper.Assert(strStr("hello", "ll") == 2)
	helper.Assert(strStr("aaaaa", "bba") == -1)
	helper.Assert(strStr("hello", "") == 0)
}
