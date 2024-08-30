package main

import (
	"leetcode/helper"
)

// 不同字符的最小子序列 同316
func smallestSubsequence(s string) string {
	count := make([]int, 256)
	for i := range s {
		count[s[i]]++
	}

	inStack := make([]bool, 256)
	stack := make([]uint8, 256)
	k := 0

	for i := range s {
		count[s[i]]--
		if inStack[s[i]] {
			continue
		}

		for k > 0 {
			top := stack[k-1]
			if s[i] >= top || count[top] < 1 {
				break
			}
			inStack[top], k = false, k-1
		}

		stack[k], k = s[i], k+1
		inStack[s[i]] = true
	}

	return string(stack[:k])
}

func main() {
	helper.Assert(smallestSubsequence("bcabc") == "abc")
	helper.Assert(smallestSubsequence("cbacdcbc") == "acdb")
}
