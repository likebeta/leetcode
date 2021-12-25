package main

import (
	"leetcode/helper"
)

// 最长有效括号
func longestValidParentheses(s string) int {
	var ans int
	if N := len(s); N > 1 {
		// dp[i]表示以i结尾最长有效括号
		dp := make([]int, N+1)
		for i := 2; i <= N; i++ {
			if s[i-1] == '(' {
				continue
			}
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				index := i - 2 - dp[i-1]
				if index >= 0 && s[index] == '(' {
					dp[i] = dp[index] + dp[i-1] + 2
				}
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}

func longestValidParentheses2(s string) int {
	var ans int
	if N := len(s); N > 1 {
		stack := []int{-1}
		for i := 0; i < N; i++ {
			if s[i] == '(' {
				stack = append(stack, i)
				continue
			}
			n := len(stack) - 1
			stack = stack[:n]
			if n == 0 {
				stack = append(stack, i)
			} else if i-stack[n-1] > ans {
				ans = i - stack[n-1]
			}
		}
	}
	return ans
}

func longestValidParentheses3(s string) int {
	var ans, left, right int
	left, right = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, 2 * right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, 2 * left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func testOne(s string, ans int) {
	helper.Assert(longestValidParentheses(s) == ans)
	helper.Assert(longestValidParentheses2(s) == ans)
	helper.Assert(longestValidParentheses3(s) == ans)
}

func main() {
	testOne("(()", 2)
	testOne("())", 2)
	testOne(")()())", 4)
	testOne("(()())", 6)
	testOne("()(()", 2)
}
