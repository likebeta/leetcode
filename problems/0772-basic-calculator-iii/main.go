package main

import (
	"leetcode/helper"
)

// 基本计算器
func calculate(s string) int {
	ans, _ := calc(s, 0)
	return ans
}

func calc(s string, i int) (int, int) {
	stk := make([]int, 0)
	sign, num := uint8('+'), 0
	var needBreak bool
	var c uint8
	for i < len(s) && !needBreak {
		c, i = s[i], i+1
		var addNum bool
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '(' {
			num, i = calc(s, i)
			addNum = true
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			addNum = true
		} else if c == ')' {
			addNum = true
			needBreak = true
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
	return sum(stk), i
}

func sum(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans += arr[i]
	}
	return ans
}

func main() {
	helper.Assert(calculate("1 + 1") == 2)
	helper.Assert(calculate(" 2-1 + 2 ") == 3)
	helper.Assert(calculate("(1+(4+5+2)-3)+(6+8)") == 23)

	helper.Assert(calculate("3+2*2") == 7)
	helper.Assert(calculate(" 3/2 ") == 1)
	helper.Assert(calculate(" 3+5 / 2 ") == 5)

	helper.Assert(calculate("6-4/2") == 4)
	helper.Assert(calculate("2*(5+5*2)/3+(6/2+8)") == 21)
}
