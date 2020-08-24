package main

import (
	"leetcode/helper"
)

// 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	helper.Assert(repeatedSubstringPattern("abab") == true)
	helper.Assert(repeatedSubstringPattern("aba") == false)
	helper.Assert(repeatedSubstringPattern("abcabcabcabc") == true)
}
