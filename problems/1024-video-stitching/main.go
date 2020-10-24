package main

import (
	"encoding/json"
	"leetcode/helper"
)

// 视频拼接
func videoStitching(clips [][]int, t int) int {
	maxn := make([]int, t)
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < t && r > maxn[l] {
			maxn[l] = r
		}
	}

	var ans, pre, last int
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return ans
}

func testOne(in string, t, ans int) {
	var arr [][]int
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		panic(err)
	}
	helper.Assert(videoStitching(arr, t) == ans)
}

func main() {
	testOne("[[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]", 10, 3)
	testOne("[[0,1],[1,2]]", 5, -1)
	testOne("[[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]]", 9, 3)
	testOne("[[0,4],[2,8]]", 5, 2)
}
