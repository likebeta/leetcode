package main

import (
	"leetcode/helper"
)

// 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {
	pre, ans := 1, 1
	for i := 2; i <= n; i++ {
		ans, pre = ans+pre, ans
	}
	return ans
}

func main() {
	helper.Log(climbStairs(2), climbStairs2(2), 2)
	helper.Log(climbStairs(3), climbStairs2(3), 3)
	helper.Log(climbStairs(4), climbStairs2(4), 5)
}
