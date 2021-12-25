package main

import "leetcode/helper"

// 正则表达式匹配
func isMatch(s string, p string) bool {
	lT, lP := len(s), len(p)
	dp := make([][]bool, lT+1, lT+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, lP+1, lP+1)
	}
	dp[lT][lP] = true
	for i := lT; i >= 0; i-- {
		for j := lP - 1; j >= 0; j-- {
			var charMatch bool
			if i < lT && (s[i] == p[j] || p[j] == '.') {
				charMatch = true
			}
			if j+1 < lP && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2]
				if !dp[i][j] && charMatch {
					dp[i][j] = dp[i+1][j]
				}
			} else if charMatch {
				dp[i][j] = dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func main() {
	helper.Assert(isMatch("aa", "a") == false)
	helper.Assert(isMatch("aa", "a*") == true)
	helper.Assert(isMatch("ab", ".*") == true)
	helper.Assert(isMatch("aab", "c*a*b") == true)
	helper.Assert(isMatch("mississippi", "mis*is*p*.") == false)
}
