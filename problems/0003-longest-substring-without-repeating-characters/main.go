package main

import (
	"leetcode/helper"
)

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	length := len(s)
	if length > 0 {
		var stat [128]int
		var start int
		for i, v := range s {
			if stat[v] > start {
				if maxLen < i-start {
					maxLen = i - start
				}
				start = stat[v]
			}
			stat[v] = i + 1
		}
		if maxLen < length-start {
			maxLen = length - start
		}
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	need := make(map[uint8]int)

	var maxLen int
	var c uint8
	var left, right int
	for right < len(s) {
		c, right = s[right], right+1
		need[c]++
		for need[c] > 1 {
			need[s[left]]--
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}

func main() {
	helper.Assert(3 == lengthOfLongestSubstring("abcabcbb"))
	helper.Assert(1 == lengthOfLongestSubstring("bbbbb"))
	helper.Assert(3 == lengthOfLongestSubstring("pwwkew"))

	helper.Assert(3 == lengthOfLongestSubstring2("abcabcbb"))
	helper.Assert(1 == lengthOfLongestSubstring2("bbbbb"))
	helper.Assert(3 == lengthOfLongestSubstring2("pwwkew"))
}
