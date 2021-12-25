package main

import (
	"leetcode/helper"
)

// 有效的山脉数组
func validMountainArray(A []int) bool {
	i, n := 0, len(A)
	for ; i+1 < n && A[i+1] > A[i]; i++ {
	}
	if i == 0 || i == n-1 {
		return false
	}
	for ; i+1 < n && A[i+1] < A[i]; i++ {
	}
	return i == n-1
}

func testOne(in string, ans bool) {
	arr := helper.ParseArray(in)
	helper.Assert(validMountainArray(arr) == ans)
}

func main() {
	testOne("[2,1]", false)
	testOne("[3,5,5]", false)
	testOne("[0,3,2,1]", true)
}
