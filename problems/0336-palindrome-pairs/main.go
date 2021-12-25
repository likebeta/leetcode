package main

import (
	"leetcode/helper"
)

// 回文对
func palindromePairs(words []string) [][]int {
	check := make(map[string]int, len(words))
	for i := range words {
		check[words[i]] = i
	}

	var ans [][]int
	for i, w := range words {
		for j := 0; j < len(w); j++ {
			if idx, ok := check[reverse(w[:j])]; ok && isPalindrome(w[j:]) && idx != i {
				ans = append(ans, []int{i, idx})
				if len(words[idx]) == 0 {
					ans = append(ans, []int{idx, i})
				}
			}
			if idx, ok := check[reverse(w[j:])]; ok && isPalindrome(w[:j]) && idx != i {
				ans = append(ans, []int{idx, i})
				if len(words[idx]) == 0 {
					ans = append(ans, []int{i, idx})
				}
			}
		}
	}
	return ans
}

func reverse(s string) string {
	res := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		res[i], res[len(s)-1-i] = res[len(s)-1-i], res[i]
	}
	return string(res)
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	helper.Log(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}), [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}})
	helper.Log(palindromePairs([]string{"bat", "tab", "cat"}), [][]int{{0, 1}, {1, 0}})
}
