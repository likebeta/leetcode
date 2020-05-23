package main

import (
	"fmt"
	"leetcode/helper"
)

// 最小覆盖子串
func minWindow(s string, t string) string {
	needMap := make(map[byte]int)
	for i := range t {
		needMap[t[i]]++
	}

	var left, right int
	needCnt := len(needMap)
	if needCnt > 0 {
		var c byte
		var i, j int
		for j < len(s) {
			c, j = s[j], j+1
			if v, ok := needMap[c]; ok {
				needMap[c]--
				if v == 1 {
					needCnt--
				}
			}
			if needCnt == 0 {
				for {
					c, i = s[i], i+1
					if v, ok := needMap[c]; ok {
						needMap[c]++
						if v == 0 {
							needCnt++
							fmt.Println(right, left, j, i)
							if right == 0 || right-left > j-i {
								left, right = i-1, j
							}
							break
						}
					}
				}
			}
		}
	}

	return s[left:right]
}

func main() {
	helper.Assert(minWindow("ADOBECODEBANC", "ABC") == "BANC")
	helper.Assert(minWindow("ADOBECODEBANC", "ABCC") == "CODEBANC")
	helper.Assert(minWindow("cabwefgewcwaefgcf", "cae") == "cwae")
	helper.Assert(minWindow("ADOBECODEBANC", "") == "")
}
