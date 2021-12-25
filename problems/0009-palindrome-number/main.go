package main

import (
	"leetcode/helper"
)

// 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var num int
	for x > num {
		num = num*10 + x%10
		x /= 10
	}
	return x == num || num/10 == x
}

func main() {
	helper.Assert(isPalindrome(0) == true)
	helper.Assert(isPalindrome(-121) == false)
	helper.Assert(isPalindrome(121) == true)
	helper.Assert(isPalindrome(10) == false)
}
