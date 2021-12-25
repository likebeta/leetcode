package main

import (
	"leetcode/helper"
)

// 回文子串
func countSubstrings(s string) int {
	var (
		n                = len(s)
		ans              int
		countPalindromic func(l, r int)
	)
	countPalindromic = func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}

	for i := 0; i < n; i++ {
		countPalindromic(i, i)
		countPalindromic(i-1, i)
	}
	return ans
}

// 归一化
func countSubstrings2(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

// dp
func countSubstrings3(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := n
	for col := 0; col < n; col++ {
		dp[col][col] = true
		for row := 0; row < col; row++ {
			if col == row + 1 {
				dp[row][col] = s[row] == s[col]
			} else {
				dp[row][col] = dp[row+1][col-1] && s[row] == s[col]
			}
			if dp[row][col] {
				ans++
			}
		}
	}
	return ans
}

func countSubstrings4(s string) int {
	n := len(s)
	dp := make([]bool, n)
	ans := n
	for col := 0; col < n; col++ {
		dp[col] = true
		for row := 0; row < col; row++ {
			if col == row + 1 {
				dp[row] = s[row] == s[col]
			} else {
				dp[row] = dp[row+1] && s[row] == s[col]
			}
			if dp[row] {
				ans++
			}
		}
	}
	return ans
}

func testOne(s string, ans int) {
	helper.Assert(countSubstrings(s) == ans)
	helper.Assert(countSubstrings2(s) == ans)
	helper.Assert(countSubstrings3(s) == ans)
	helper.Assert(countSubstrings4(s) == ans)
}

func main() {
	testOne("abc", 3)
	testOne("aaa", 6)
	testOne("abcba", 7)
	testOne("fdsklf", 6)
}
