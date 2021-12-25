package main

import (
	"leetcode/helper"
)

// 电话号码的字母组合
func letterCombinations(digits string) []string {
	var (
		phoneMap = map[byte]string{
			'2': "abc",
			'3': "def",
			'4': "ghi",
			'5': "jkl",
			'6': "mno",
			'7': "pqrs",
			'8': "tuv",
			'9': "wxyz",
		}
		combinations []string
		backtrack    func(int, []byte)
		N            = len(digits)
	)

	backtrack = func(index int, combination []byte) {
		if index == N {
			combinations = append(combinations, string(combination))
		} else {
			letters := phoneMap[digits[index]]
			lettersCount := len(letters)
			for i := 0; i < lettersCount; i++ {
				combination[index] = letters[i]
				backtrack(index+1, combination)
			}
		}
	}

	if N > 0 {
		combination := make([]byte, N)
		backtrack(0, combination)
	}
	return combinations
}

func main() {
	helper.Log(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
}
