package main

import (
	"leetcode/helper"
)

// 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	need := make(map[uint8]int)
	for i := range s1 {
		need[s1[i]]++
	}

	windows := make(map[uint8]int)
	needCnt := len(need)

	var c uint8
	var left, right int
	for right < n {
		c, right = s2[right], right+1
		if _, ok := need[c]; !ok {
			left, needCnt = right, len(need)
			clear(windows)
			continue
		}

		windows[c]++
		if windows[c] == need[c] {
			needCnt--
		}

		if right-left == m {
			if needCnt == 0 {
				return true
			}
			c, left = s2[left], left+1
			if windows[c] == need[c] {
				needCnt++
			}
			windows[c]--
		}
	}

	return false
}

func checkInclusion2(s1, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	need := make(map[uint8]int)
	for i := range s1 {
		need[s1[i]]++
	}

	var c uint8
	var left, right int
	for right < n {
		c, right = s2[right], right+1
		need[c]--
		for need[c] < 0 {
			need[s2[left]]++
			left++
		}
		if right-left == m {
			return true
		}
	}
	return false
}

func main() {
	helper.Assert(checkInclusion("ab", "eidbaooo") == true)
	helper.Assert(checkInclusion2("ab", "eidbaooo") == true)
	helper.Assert(checkInclusion("ab", "eidboaoo") == false)
	helper.Assert(checkInclusion2("ab", "eidboaoo") == false)
	helper.Assert(checkInclusion("adc", "dcda") == true)
	helper.Assert(checkInclusion2("adc", "dcda") == true)
}
