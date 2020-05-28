package main

import (
	"leetcode/helper"
	"strings"
)

// 字符串解码
func decodeString(s string) string {
	var (
		sLen        = len(s)
		index       int
		parse       func() string
		parseNumber func() int
	)

	parseNumber = func() int {
		num := 0
		for s[index] >= '0' && s[index] <= '9' {
			num = num*10 + int(s[index]-'0')
			index++
		}
		return num
	}

	parse = func() string {
		var ans string
		var num, start int
		for start = index; index < sLen; index++ {
			if s[index] == ']' {
				break
			} else if s[index] >= '0' && s[index] <= '9' {
				ans += s[start:index]
				num = parseNumber()
				index++
				ans += strings.Repeat(parse(), num)
				start = index + 1
			}
		}
		return ans + s[start:index]
	}

	return parse()
}

func main() {
	helper.Log(decodeString("cc3[a]2[bc]"), "ccaaabcbc")
	helper.Log(decodeString("3[a2[c]]"), "accaccacc")
	helper.Log(decodeString("2[abc]3[cd]ef"), "abcabccdcdcdef")
}
