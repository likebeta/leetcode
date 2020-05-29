package main

import (
	"leetcode/helper"
	"strings"
)

// 复原IP地址
func restoreIpAddresses(s string) []string {
	var (
		sLen    = len(s)
		ans     []string
		ip      = make([]string, 4, 4)
		process func(int, int)
	)
	process = func(i int, part int) {
		if i == sLen && part == 0 {
			ans = append(ans, strings.Join(ip, "."))
			return
		}
		if i == sLen || part == 0 {
			return
		}
		if i+part <= sLen {
			ip[4-part] = s[i : i+1]
			process(i+1, part-1)
		}
		if s[i] != '0' {
			if i+1+part <= sLen {
				ip[4-part] = s[i : i+2]
				process(i+2, part-1)
			}
			if i+2+part <= sLen {
				num := int(s[i] - '0')
				num = num*10 + int(s[i+1]-'0')
				num = num*10 + int(s[i+2]-'0')
				if num <= 255 {
					ip[4-part] = s[i : i+3]
					process(i+3, part-1)
				}
			}
		}
	}
	process(0, 4)
	return ans
}

func main() {
	helper.Log(restoreIpAddresses("25525511135"))
	helper.Log(restoreIpAddresses("010010"))
	helper.Log(restoreIpAddresses("000256"))
}
