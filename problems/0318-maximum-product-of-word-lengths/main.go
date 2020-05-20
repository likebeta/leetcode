package main

import (
	"leetcode/helper"
)

// 最大单词长度乘积
func maxProduct(words []string) int {
	cache := make(map[int]int)
	for i := range words {
		var mask int
		for j := range words[i] {
			mask |= 1 << (words[i][j] - 'a')
		}
		if cache[mask] < len(words[i]) {
			cache[mask] = len(words[i])
		}
	}

	var ans int
	for k1, v1 := range cache {
		for k2, v2 := range cache {
			if k1&k2 == 0 && v1*v2 > ans {
				ans = v1 * v2
			}
		}
	}
	return ans
}

func main() {
	helper.Assert(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}) == 16)
	helper.Assert(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}) == 4)
	helper.Assert(maxProduct([]string{"a", "aa", "aaa", "aaaa"}) == 0)
}
