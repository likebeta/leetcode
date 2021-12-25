package main

import (
	"leetcode/helper"
	"sort"
)

// 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}

func testOne(in string, ans int) {
	var points [][]int
	helper.Load([]byte(in), &points)
	helper.Assert(findMinArrowShots(points) == ans)
}

func main() {
	testOne("[[10,16],[2,8],[1,6],[7,12]]", 2)
	testOne("[[1,2],[3,4],[5,6],[7,8]]", 4)
	testOne("[[1,2],[2,3],[3,4],[4,5]]", 2)
	testOne("[[1,2]]", 1)
	testOne("[[2,3],[2,3]]", 1)
	testOne("[[-2147483646,-2147483645],[2147483646,2147483647]]", 2)
}
