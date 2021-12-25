package main

import (
	"leetcode/helper"
)

// 判断子序列
func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(s)
}

func main() {
	helper.Assert(isSubsequence("abc", "ahbgdc") == true)
	helper.Assert(isSubsequence("axc", "ahbgdc") == false)
}
