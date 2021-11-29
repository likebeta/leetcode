package main

import (
	"leetcode/helper"
)

// 每日温度
func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := make([]int, len(T)+1)
	var k int
	for i := range T {
		for top := stack[k]; k != 0 && T[top] < T[i]; top = stack[k] {
			ans[top], k = i-top, k-1
		}
		k++
		stack[k] = i
	}
	return ans
}

func testOne(in string, ans string) {
	nums := helper.ParseArray(in)
	result := dailyTemperatures(nums)
	helper.Assert(helper.DumpArray(result) == ans)
	helper.Log(in, "=>", ans)
}

func main() {
	testOne("[73,74,75,71,69,72,76,73]", "[1,1,4,2,1,1,0,0]")
	testOne("[30,40,50,60]", "[1,1,1,0]")
	testOne("[30,60,90]", "[1,1,0]")
}
