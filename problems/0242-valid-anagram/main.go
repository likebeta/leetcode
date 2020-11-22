package main

import (
	"leetcode/helper"
)

// 有效的字母异位词
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func testOne(s, t string, ans bool) {
	helper.Assert(isAnagram(s, t) == ans)
}

func main() {
	testOne("anagram", "nagaram", true)
	testOne("rat", "car", false)
}
