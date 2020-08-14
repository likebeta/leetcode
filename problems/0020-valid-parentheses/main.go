package main

import (
	"leetcode/helper"
)

// 有效的括号
func isValid(s string) bool {
	var count int
	if length := len(s); length > 0 {
		if length%2 != 0 {
			return false
		}
		bracketsMap := map[byte]byte{']': '[', ')': '(', '}': '{'}
		half := length / 2
		stack := make([]byte, half)
		for i := range s {
			if v, ok := bracketsMap[s[i]]; !ok {
				if count >= half {
					return false
				}
				stack[count] = s[i]
				count++
			} else {
				if count == 0 || stack[count-1] != v {
					return false
				}
				count--
			}
		}
	}

	return count == 0
}

func main() {
	helper.Assert(isValid("()") == true)
	helper.Assert(isValid("()[]{}") == true)
	helper.Assert(isValid("(]") == false)
	helper.Assert(isValid("([)]") == false)
	helper.Assert(isValid("{[]}") == true)
}
