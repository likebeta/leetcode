package main

import (
	"leetcode/helper"
)

// 有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	cnt := make([]int, 101)
	for _, v := range nums {
		cnt[v]++
	}
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}

func testOne(in string, ans string) {
	arr := helper.ParseArray(in)
	result := smallerNumbersThanCurrent(arr)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("[8,1,2,2,3]", "[4,0,1,1,3]")
	testOne("[6,5,4,8]", "[2,1,0,3]")
	testOne("[7,7,7,7]", "[0,0,0,0]")
}
