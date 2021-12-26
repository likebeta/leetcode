package main

import (
	"leetcode/helper"
	"strings"
)

// Bigram 分词
func findOcurrences(text, first, second string) (ans []string) {
	words := strings.Split(text, " ")
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return
}

func testOne(text, first, second string, ans string) {
	result := findOcurrences(text, first, second)
	helper.Assert(helper.Dump(result) == ans)
}

func main() {
	testOne("alice is a good girl she is a good student", "a", "good", `["girl","student"]`)
	testOne("we will we will rock you", "we", "will", `["we","rock"]`)
}
