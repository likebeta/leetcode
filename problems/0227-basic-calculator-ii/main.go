package main

import (
	"leetcode/helper"
)

// 基本计算器 II
func calculate(s string) int {
	stk := make([]int, 0)
	sign, num := uint8('+'), 0
	var c uint8
	var i int
	for i < len(s) {
		c, i = s[i], i+1
		var addNum bool
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			addNum = true
		}

		if i == len(s) {
			addNum = true
		}

		if addNum {
			switch sign {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			case '*':
				stk[len(stk)-1] *= num
			case '/':
				stk[len(stk)-1] /= num
			}
			sign, num = c, 0
		}
	}
	return sum(stk)
}

func sum(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans += arr[i]
	}
	return ans
}

func main() {
	helper.Assert(calculate("3+2*2") == 7)
	helper.Assert(calculate(" 3/2 ") == 1)
	helper.Assert(calculate(" 3+5 / 2 ") == 5)
}
