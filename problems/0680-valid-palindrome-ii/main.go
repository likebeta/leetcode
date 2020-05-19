package main

import (
	"leetcode/helper"
)

// 验证回文字符串 Ⅱ
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j || isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
}

func isPalindrome(s string, i, j int) bool {
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j
}

func main() {
	helper.Assert(validPalindrome("aba") == true)
	helper.Assert(validPalindrome("abca") == true)
}
