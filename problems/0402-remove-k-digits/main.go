package main

import (
	"leetcode/helper"
	"strings"
)

// 移掉K位数字
func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}


func main() {
	helper.Assert(removeKdigits("1432219", 3) == "1219")
	helper.Assert(removeKdigits("10200", 1) == "200")
	helper.Assert(removeKdigits("10", 2) == "0")
}
