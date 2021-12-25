package main

import (
	"leetcode/helper"
)

// 插入区间
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testOne(in1 string, in2 string, ans string) {
	var intervals [][]int
	helper.Load([]byte(in1), &intervals)
	var newInterval []int
	helper.Load([]byte(in2), &newInterval)
	result := insert(intervals, newInterval)
	helper.Log(helper.Dump(result), ans)
}

func main() {
	testOne("[[1,3],[6,9]]", "[2,5]", "[[1,5],[6,9]]")
	testOne("[[1,2],[3,5],[6,7],[8,10],[12,16]]", "[4,8]", "[[1,2],[3,10],[12,16]]")
}
