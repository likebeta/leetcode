package main

import (
	"leetcode/helper"
)

// 比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	for i >= 0 || j >= 0 {
		for skip := 0; i >= 0; i-- {
			if S[i] == '#' {
				skip++
			} else if skip > 0 {
				skip--
			} else {
				break
			}
		}

		for skip := 0; j >= 0; j-- {
			if T[j] == '#' {
				skip++
			} else if skip > 0{
				skip--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		} else {
			return true
		}
		i--
		j--
	}
	return true
}

func main() {
	helper.Assert(backspaceCompare("ab#c", "ad#c"))
	helper.Assert(backspaceCompare("ab##", "c#d#"))
	helper.Assert(backspaceCompare("a##c", "#a#c"))
	helper.Assert(!backspaceCompare("a#c", "b"))
}
