package main

import (
	"leetcode/helper"
)

// 字符串相加
func addStrings(num1 string, num2 string) string {
	N, M := len(num1), len(num2)
	if N < M {
		return addStrings(num2, num1)
	}

	ans := make([]byte, N+1)
	var carryFlag, tmp byte
	for N--; N >= 0; N-- {
		tmp = carryFlag + (num1[N] - '0')
		if M > 0 {
			M--
			tmp += num2[M] - '0'
		}
		ans[N+1] = tmp % 10 + '0'
		carryFlag = tmp / 10
	}
	if carryFlag > 0 {
		ans[0] = carryFlag + '0'
	} else {
		ans = ans[1:]
	}
	return string(ans)
}

func main() {
	helper.Log(addStrings("0", "0"), "0")
	helper.Log(addStrings("9", "1"), "10")
}
