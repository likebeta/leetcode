package main

import (
	"leetcode/helper"
)

// 反转字符串中的单词 III
func reverseWords(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	b := make([]byte, length)
	var start, end int
	for i := 0; i < length; i++ {
		if s[i] == ' ' || i == length-1 {
			if i == length-1 {
				end = i
			} else {
				end = i - 1
			}
			for start <= end {
				b[start], b[end] = s[end], s[start]
				start++
				end--
			}
			if s[i] == ' ' {
				b[i] = ' '
				start = i + 1
			}
		}
	}
	return string(b)
}

func main() {
	helper.Assert(reverseWords("Let's take LeetCode contest") == "s'teL ekat edoCteeL tsetnoc")
}
