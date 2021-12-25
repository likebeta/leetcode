package main

import (
	"leetcode/helper"
)

// 最长回文子串
func longestPalindrome(s string) string {
	length := len(s)
	left, right := -1, 0
	for i := range s {
		l1, r1 := extend(s, length, i, i)
		if r1-l1 > right-left {
			left, right = l1, r1
		}
		l2, r2 := extend(s, length, i, i+1)
		if r2-l2 > right-left {
			left, right = l2, r2
		}
	}
	return s[left+1 : right]
}

func extend(s string, length, left, right int) (int, int) {
	for left > -1 && right < length && s[left] == s[right] {
		left--
		right++
	}
	return left, right
}

func main() {
	helper.Assert(longestPalindrome("babad") == "bab")
	helper.Assert(longestPalindrome("cbbd") == "bb")
}
