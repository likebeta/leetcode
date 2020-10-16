package main

import (
	"leetcode/helper"
)

// 有序数组的平方
func sortedSquares(A []int) []int {
	var ans []int
	if N := len(A); N > 0 {
		ans = make([]int, N)
		left, right := 0, N-1
		for i := N - 1; i >= 0; i-- {
			if v, w := A[left]*A[left], A[right]*A[right]; v > w {
				ans[i] = v
				left++
			} else {
				ans[i] = w
				right--
			}
		}
	}
	return ans
}

func testOne(in string, ans string) {
	arr := helper.ParseArray(in)
	result := sortedSquares(arr)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("[-4,-1,0,3,10]", "[0,1,9,16,100]")
	testOne("[-7,-3,2,3,11]", "[4,9,9,49,121]")
}
