package main

import (
	"leetcode/helper"
)

// 斐波那契数
func fib(N int) int {
	if N < 2 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func fib2(N int) int {
	if N < 2 {
		return N
	}
	ans, prePre, pre := 0, 0, 1
	for i := 2; i <= N; i++ {
		ans = pre + prePre
		pre, prePre = ans, pre
	}
	return ans
}

func main() {
	helper.Log(fib(2), fib2(2), 1)
	helper.Log(fib(3), fib2(3), 2)
	helper.Log(fib(4), fib2(4), 3)
}
