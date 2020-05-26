package main

import (
	"leetcode/helper"
)

// 括号生成
func generateParenthesis(n int) []string {
	var ans []string
	if n > 0 {
		arr := make([]byte, n<<1)
		var process func([]byte, int, int, int)
		process = func(bytes []byte, n, l, r int) {
			if l == n && r == n {
				ans = append(ans, string(arr))
				return
			}
			if l < n {
				arr[l+r] = '('
				process(arr, n, l+1, r)
			}

			if r < l {
				arr[l+r] = ')'
				process(arr, n, l, r+1)
			}
		}
		process(arr, n, 0, 0)
	}
	return ans
}

func main() {
	helper.Log(generateParenthesis(1))
	helper.Log(generateParenthesis(2))
	helper.Log(generateParenthesis(3))
}
