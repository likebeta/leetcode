package main

import (
	"leetcode/helper"
)

// 拼车
func carPooling(trips [][]int, capacity int) bool {
	ans := make([]int, 1024)
	for _, one := range trips {
		num, i, j := one[0], one[1], one[2]
		ans[i] += num
		ans[j] -= num
	}

	n := 0
	for i := 0; i < 1024; i++ {
		n += ans[i]
		if n > capacity {
			return false
		}
	}

	return true
}

func testOne(capacity int, updates string, ans bool) {
	matrix := helper.ParseMatrix[int](updates)
	ret := carPooling(matrix, capacity)
	helper.Assert(ret == ans)
}

func main() {
	testOne(4, "[[2,1,5],[3,3,7]]", false)
	testOne(5, "[[2,1,5],[3,3,7]]", true)
}
