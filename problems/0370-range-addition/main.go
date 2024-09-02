package main

import (
	"leetcode/helper"
)

// 区间加法
func getModifiedArray(length int, updates [][]int) []int {
	ans := make([]int, length)
	for _, one := range updates {
		i, j, add := one[0], one[1], one[2]
		ans[i] += add
		if j+1 < length {
			ans[j+1] -= add
		}
	}

	for i := 1; i < length; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

func testOne(length int, updates string, ans string) {
	matrix := helper.ParseMatrix[int](updates)
	ret := getModifiedArray(length, matrix)
	helper.Assert(helper.DumpArray(ret) == ans)
}

func main() {
	testOne(5, "[[1,3,2],[2,4,3],[0,2,-2]]", "[-2,0,3,5,3]")
}
