package main

import (
	"leetcode/helper"
)

// 平衡括号字符串的最少插入次数
func minInsertions(s string) int {
	var need, ans int

	for _, c := range s {
		if c == '(' {
			if need%2 == 1 { // 奇数)+(, 补一个); need+2-1=need+1
				ans++
				need += 1
			} else {
				need += 2
			}
		} else {
			if need == 0 { // 多一个), 补一个(; need-1+2=0+1=1
				ans++
				need = 1
			} else {
				need--
			}
		}
	}

	return ans + need
}

func main() {
	helper.Assert(minInsertions("(()))") == 1)
	helper.Assert(minInsertions("())") == 0)
	helper.Assert(minInsertions("))())(") == 3)
	helper.Assert(minInsertions("((((((") == 12)
	helper.Assert(minInsertions(")))))))") == 5)
}
