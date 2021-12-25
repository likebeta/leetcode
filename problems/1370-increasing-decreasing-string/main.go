package main

import (
	"leetcode/helper"
)

// 上升下降字符串
func sortString(s string) string {
	cnt := ['z' + 1]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) < n {
		for i := byte('a'); i <= 'z'; i++ {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
		for i := byte('z'); i >= 'a'; i-- {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
	}
	return string(ans)
}

func testOne(s string, ans string) {
	helper.Assert(sortString(s) == ans)
}

func main() {
	testOne("aaaabbbbcccc", "abccbaabccba")
	testOne("rat", "art")
	testOne("leetcode", "cdelotee")
	testOne("ggggggg", "ggggggg")
	testOne("spo", "ops")
}
