package main

import (
	"leetcode/helper"
)

// 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return nil
	}

	need := make(map[uint8]int)
	for i := range p {
		need[p[i]]++
	}

	var ans []int
	var c uint8
	var left, right int
	for right < n {
		c, right = s[right], right+1
		need[c]--
		for need[c] < 0 {
			need[s[left]]++
			left++
		}
		if right-left == m {
			ans = append(ans, left)
		}
	}
	return ans
}

func testOne(s, p string, out string) {
	ans := findAnagrams(s, p)
	sAns := helper.DumpArray(ans)
	helper.Assert(sAns == out)
}

func main() {
	testOne("cbaebabacd", "abc", "[0,6]")
	testOne("abab", "ab", "[0,1,2]")
}
