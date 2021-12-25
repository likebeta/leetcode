package main

import (
	"leetcode/helper"
)

// 数字范围按位与
func rangeBitwiseAnd(m int, n int) int {
	ans := m
	for m++; m <= n; m++ {
		ans &= m
	}
	return ans
}

func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n &= n - 1
	}
	return n
}

func testOne(m, n int, ans int) {
	helper.Assert(rangeBitwiseAnd(m, n) == ans)
	helper.Assert(rangeBitwiseAnd2(m, n) == ans)
}

func main() {
	testOne(5, 7, 4)
	testOne(0, 1, 0)
}
