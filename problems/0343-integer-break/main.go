package main

import (
	"leetcode/helper"
	"math"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for k := 1; k <= i/2; k++ {
			dp[i] = max3(dp[i], k*(i-k), k*dp[i-k])
		}
	}
	return dp[n]
}

func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = max4(2*(i-2), 2*dp[i-2], 3*(i-3), 3*dp[i-3])
	}
	return dp[n]
}

func max3(a, b, c int) int {
	if a < b {
		a = b
	}
	if a < c {
		return c
	}
	return a
}

func max4(a, b, c, d int) int {
	if a < b {
		a = b
	}
	if a < c {
		a = c
	}
	if a < d {
		return d
	}
	return a
}

func integerBreak3(n int) int {
	if n <= 3 {
		return n - 1
	}
	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(quotient)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(quotient-1))) * 4
	}
	return int(math.Pow(3, float64(quotient))) * 2
}

func testOne(n int, ans int) {
	helper.Assert(integerBreak(n) == ans)
	helper.Assert(integerBreak2(n) == ans)
	helper.Assert(integerBreak3(n) == ans)
}

func main() {
	testOne(2, 1)
	testOne(3, 2)
	testOne(4, 4)
	testOne(5, 6)
	testOne(6, 9)
	testOne(7, 12)
	testOne(8, 18)
	testOne(9, 27)
	testOne(10, 36)
}
