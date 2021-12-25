package main

import "leetcode/helper"

// 二进制求和
func addBinary(a string, b string) string {
	var ans string
	var carry int
	lenA, lenB := len(a), len(b)
	max := lenA
	if lenB > lenA {
		max = lenB
	}

	for i := 0; i < max; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		if carry%2 == 1 {
			ans = "1" + ans
		} else {
			ans = "0" + ans
		}
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func main() {
	helper.Assert(addBinary("11", "1") == "100")
	helper.Assert(addBinary("1010", "1011") == "10101")
}
