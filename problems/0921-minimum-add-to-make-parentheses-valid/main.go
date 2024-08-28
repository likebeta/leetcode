package main

import (
	"leetcode/helper"
)

// 使括号有效的最少添加
func minAddToMakeValid(s string) int {
	var need, ans int

	for _, c := range s {
		if c == '(' {
			need++
		} else {
			if need == 0 {
				ans++
			} else {
				need--
			}
		}
	}

	return need + ans
}

func main() {
	helper.Assert(minAddToMakeValid("())") == 1)
	helper.Assert(minAddToMakeValid("(((") == 3)
}
