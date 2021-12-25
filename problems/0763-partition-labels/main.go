package main

import (
	"leetcode/helper"
)

// 划分字母区间
func partitionLabels(S string) []int {
	lastPos := make([]int, 26)
	for i, c := range S {
		lastPos[c-'a'] = i
	}
	var ans []int
	start, end := 0, 0
	for i, c := range S {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		} else if i == end {
			ans = append(ans, end-start+1)
			start, end = i+1, i+1
		}
	}
	return ans
}

func testOne(in string, ans string) {
	result := partitionLabels(in)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("ababcbacadefegdehijhklij", "[9,7,8]")
	testOne("eaaaabaaec", "[9,1]")
}
