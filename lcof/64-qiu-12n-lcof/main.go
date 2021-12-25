package main

import "leetcode/helper"

// 求1+2+…+n
func sumNums(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}

func main() {
	helper.Assert(sumNums(3) == 6)
	helper.Assert(sumNums(9) == 45)
}

