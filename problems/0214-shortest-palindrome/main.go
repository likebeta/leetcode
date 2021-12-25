package main

import (
	"leetcode/helper"
)

// 最短回文串
func shortestPalindrome(s string) string {
	if N := len(s); N > 1 {
		arr := make([]byte, N)
		for i := 0; i < N; i++ {
			arr[i] = s[N-1-i]
		}
		rs := string(arr)
		index := search(rs, s)
		if index != N {
			s = rs[0:N-index] + s
		}
	}
	return s
}

func getNext(str string) []int {
	length := len(str)
	if length == 0 {
		return nil
	}
	nextArr := make([]int, length, length)
	nextArr[0] = -1
	cn := 0
	for pos := 2; pos < length; {
		if str[pos-1] == str[cn] {
			cn++
			nextArr[pos] = cn
			pos++
		} else if nextArr[cn] == -1 {
			nextArr[pos] = 0
			pos++
		} else {
			cn = nextArr[cn]
		}
	}
	return nextArr
}

func search(src string, pattern string) int {
	nextArr := getNext(pattern)
	var i, j int
	for i < len(src) && j < len(pattern) {
		if src[i] == pattern[j] {
			i++
			j++
		} else if nextArr[j] == -1 {
			i++
		} else {
			j = nextArr[j]
		}
	}
	return j
}

func main() {
	helper.Assert(shortestPalindrome("aacecaaa") == "aaacecaaa")
	helper.Assert(shortestPalindrome("abcd") == "dcbabcd")
	helper.Assert(shortestPalindrome("aaa") == "aaa")
}
