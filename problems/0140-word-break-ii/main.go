package main

import (
	"leetcode/helper"
	"strings"
)

// 单词拆分2
func wordBreak(s string, wordDict []string) []string {
	var (
		n         = len(s)
		wordSet   = make(map[string]bool)
		dp        = make([][][]string, n)
		backtrack func(index int) [][]string
		ans       = make([]string, 0)
	)

	for _, w := range wordDict {
		wordSet[w] = true
	}

	if !canWordBreak(s, wordSet) {
		return ans
	}

	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		var wordsList [][]string
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if wordSet[word] {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if wordSet[word] {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		ans = append(ans, strings.Join(words, " "))
	}
	return ans
}

func canWordBreak(s string, wordSet map[string]bool) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func testOne(s, in string, ans string) {
	var wordDict []string
	helper.Load([]byte(in), &wordDict)
	result := wordBreak(s, wordDict)
	helper.Log(helper.Dump(result), ans)
}

func main() {
	testOne("catsanddog", `["cat", "cats", "and", "sand", "dog"]`, `["cats and dog","cat sand dog"]`)
	testOne("pineapplepenapple", `["apple", "pen", "applepen", "pine", "pineapple"]`, `["pine apple pen apple","pineapple pen apple","pine applepen apple"]`)
	testOne("catsandog", `["cats", "dog", "sand", "and", "cat"]`, `[]`)
}
