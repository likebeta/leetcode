package main

import "leetcode/helper"

// 正则表达式匹配
// dp[i][j] 表示字符串s[i:]和模式p[j:]是否匹配（即s从i开始到末尾的子串和p从j开始到末尾的子串是否匹配）
// 状态转移方程分为以下几种情况：
// 1. 当前字符匹配（s[i] == p[j] 或 p[j] == '.'）
//    - 如果p[j+1]是'*'：可以选择匹配0次或多次
//      * 匹配0次：dp[i][j] = dp[i][j+2]（跳过当前字符和'*'，尝试匹配s[i:]和p[j+2:]）
//      * 匹配多次：dp[i][j] = dp[i+1][j]（当前匹配成功，继续用同一个模式匹配s[i+1:]和p[j:]）
//    - 如果p[j+1]不是'*'：dp[i][j] = dp[i+1][j+1]（当前匹配，继续匹配s[i+1:]和p[j+1:]）
// 2. 当前字符不匹配
//    - 如果p[j+1]是'*'：dp[i][j] = dp[i][j+2]（只能选择匹配0次，跳过当前字符和'*'）
//    - 如果p[j+1]不是'*'：dp[i][j] = false（当前字符不匹配且无法跳过，整体无法匹配）
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// 创建二维DP数组，dp[i][j]表示s[i:]和p[j:]是否匹配
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// 空字符串匹配空模式
	dp[m][n] = true
	for j := n - 2; j >= 0; j-- {
		if p[j+1] == '*' {
			dp[m][j] = dp[m][j+2]
		}
	}

	// 从后向前遍历，i从m开始处理空字符串的情况
	for i := m - 1; i >= 0; i-- {
		// j从n-1开始，因为当p为空时只有s也为空才匹配
		for j := n - 1; j >= 0; j-- {
			// 当前字符匹配
			if p[j] == s[i] || p[j] == '.' {
				if j+1 < n && p[j+1] == '*' {
					// 如果p[j+1]是'*'：可以选择匹配0次或多次
					// 匹配0次：dp[i][j+2]（跳过当前字符和'*'）
					// 匹配多次：dp[i+1][j]（当前匹配成功，继续用同一个模式）
					dp[i][j] = dp[i][j+2] || dp[i+1][j]
				} else {
					// 如果p[j+1]不是'*'：当前匹配，继续匹配后续子串
					dp[i][j] = dp[i+1][j+1]
				}
			} else {
				// 检查下一个字符是否为'*'
				if j+1 < n && p[j+1] == '*' {
					// 如果p[j+1]是'*'：只能选择匹配0次，跳过当前字符和'*'
					dp[i][j] = dp[i][j+2]
				} else {
					// 如果p[j+1]不是'*'：当前字符不匹配且无法跳过，整体无法匹配
					dp[i][j] = false
				}
			}
		}
	}

	return dp[0][0]
}

// isMatch2 从前往后遍历实现正则表达式匹配
// dp[i][j] 表示s的前i个字符和p的前j个字符是否匹配
func isMatch2(s string, p string) bool {
	lT, lP := len(s), len(p)
	// 创建dp数组
	dp := make([][]bool, lT+1)
	for i := 0; i <= lT; i++ {
		dp[i] = make([]bool, lP+1)
	}

	// 初始化：空字符串和空模式匹配
	dp[0][0] = true

	// 处理模式以*开头的特殊情况
	for j := 1; j <= lP; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 因为从前往后遍历, 每次循环是为了处理s[0:i]和p[0:j]的匹配结果,
	// 不要去想s[0:i+1]或p[0:j+1]怎么匹配，否则容易陷进去
	for i := 1; i <= lT; i++ {
		for j := 1; j <= lP; j++ {
			// 当前字符是'*'
			if p[j-1] == '*' {
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					// 和当前字符匹配，可以选择匹配0次或1次(匹配多次的情况是通过多次选择匹配叠加而成)
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					// 和当前字符不匹配，只能让'p[j-1]*'匹配0次，取s[0:i]和p[0:j-2]的匹配结果(dp[i][j-2])
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if p[j-1] == '.' || p[j-1] == s[i-1] {
					// 当前字符匹配
					dp[i][j] = dp[i-1][j-1]
				} /* else {
					// 当前字符不匹配
					dp[i][j] = false
				} */
			}
		}
	}

	return dp[lT][lP]
}

func testOne(s, p string, ans bool) {
	helper.Assert(isMatch(s, p) == ans)
	helper.Assert(isMatch2(s, p) == ans)
}

func main() {
	testOne("aa", "a", false)
	testOne("aa", "a*", true)
	testOne("ab", ".*", true)
	testOne("aab", "c*a*b", true)
	testOne("mississippi", "mis*is*p*.", false)
}
