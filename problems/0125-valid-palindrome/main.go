package main

import (
	"leetcode/helper"
)

// 验证回文串
func isPalindrome(s string) bool {
	const DIFF = 'a' - 'A'
	i, j := 0, len(s)-1
	var ti, tj uint8
	for i < j {
		if s[i] != s[j] {
			if ti = charType(s[i]); ti == 0 {
				i++
				continue
			}
			if tj = charType(s[j]); tj == 0 {
				j--
				continue
			}
			if ti != tj || (s[i]+DIFF != s[j] && s[i] != s[j]+DIFF) {
				return false
			}
		}
		i++
		j--
	}
	return i >= j
}

func charType(c byte) uint8 {
	if c >= '0' && c <= '9' {
		return 1
	}
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return 2
	}
	return 0
}

func main() {
	helper.Log(isPalindrome("A man, a plan, a canal: Panama") == true)
	helper.Log(isPalindrome("race a car") == false)
}
