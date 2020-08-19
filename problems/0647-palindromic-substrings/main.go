package main

import (
	"leetcode/helper"
)

// 回文子串
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func main() {
	helper.Assert(countSubstrings("abc") == 3)
	helper.Assert(countSubstrings("aaa") == 6)
}
