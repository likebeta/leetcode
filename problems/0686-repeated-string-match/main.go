package main

import (
	"leetcode/helper"
)

// 重复叠加字符串匹配
func repeatedStringMatch(a string, b string) int {
	var checkIndexes []int
	var stat [26]bool
	for i := range a {
		stat[a[i]-'a'] = true
		if a[i] == b[0] {
			checkIndexes = append(checkIndexes, i)
		}
	}

	ans := -1
	for i := range checkIndexes {
		success, index := stringMatch(a, b, checkIndexes[i])
		if success {
			ans = (checkIndexes[i] + len(b) + len(a) - 1) / len(a)
			break
		}

		if !stat[b[index]-'a'] {
			break
		}
	}
	return ans
}

func stringMatch(a, b string, start int) (bool, int) {
	for i := range b {
		if a[start%len(a)] != b[i] {
			return false, i
		}
		start++
	}
	return true, len(b)
}

func testOne(a, b string, ans int) {
	result := repeatedStringMatch(a, b)
	helper.LogF("[%s, %s] => %d", a, b, result)
	helper.Assert(result == ans)
}

func main() {
	testOne("abcd", "cdabcdab", 3)
	testOne("a", "aa", 2)
	testOne("a", "a", 1)
	testOne("abc", "wxyz", -1)
	testOne("aa", "a", 1)
	testOne("aaac", "aac", 1)
	testOne("abc", "cabcabca", 4)
}
