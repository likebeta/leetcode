package main

import (
	"leetcode/helper"
)

// 把数字翻译成字符串
func translateNum(num int) int {
	ans, dp1, dp2, unit1 := 1, 1, 1, num%10
	for num /= 10; num != 0; num /= 10 {
		ans = dp1
		unit := num % 10
		if unit != 0 && (unit < 2 || (unit == 2 && unit1 < 6)) {
			ans += dp2
		}
		dp1, dp2, unit1 = ans, dp1, unit
	}
	return ans
}

func translateNum1(num int) int {
	ans, dp1, dp2 := 1, 1, 0
	for unit1 := 0; num != 0; num /= 10 {
		ans = dp1
		unit := num % 10
		if unit != 0 && (unit < 2 || (unit == 2 && unit1 < 6)) {
			ans += dp2
		}
		dp1, dp2, unit1 = ans, dp1, unit
	}
	return ans
}

func main() {
	helper.Log(translateNum1(0), 1)
	helper.Log(translateNum1(4), 1)
	helper.Log(translateNum1(12), 2)
	helper.Log(translateNum1(126), 2)
	helper.Log(translateNum1(624), 2)
	helper.Log(translateNum1(12258), 5)
}
