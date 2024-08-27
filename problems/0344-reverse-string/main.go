package main

import (
	"leetcode/helper"
)

// 反转字符串
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func testOne(in, out string) {
	inBytes := []byte(in)
	reverseString(inBytes)
	helper.Assert(string(inBytes) == out)
}

func main() {
	testOne("Hello", "olleH")
	testOne("Hannah", "hannaH")
}
