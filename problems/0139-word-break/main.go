package main

import "leetcode/helper"

// 单词拆分
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[length]
}

func main() {
	helper.Assert(wordBreak("leetcode", []string{"leet", "code"}) == true)
	helper.Assert(wordBreak("applepenapple", []string{"apple", "pen"}) == true)
	helper.Assert(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)
}
