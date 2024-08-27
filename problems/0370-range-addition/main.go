package main

import (
	"leetcode/helper"
)

// 区间加法
func getModifiedArray(length int, updates [][]int) []int {
	d := make([]int, length)
	for _, one := range updates {
		i, j, add := one[0], one[1], one[2]
		d[i] += add
		if j+1 < length {
			d[j+1] -= add
		}
	}

	for i := 1; i < length; i++ {
		d[i] += d[i-1]
	}
	return d
}

func testOne(length int, updates string, ans string) {
	matrix := helper.ParseMatrix[int](updates)
	ret := getModifiedArray(length, matrix)
	helper.Assert(helper.DumpArray[int](ret) == ans)
}

func main() {
	testOne(5, "[[1,3,2],[2,4,3],[0,2,-2]]", "[-2,0,3,5,3]")
}
