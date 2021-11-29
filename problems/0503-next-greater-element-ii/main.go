package main

import (
	"leetcode/helper"
)

// 下一个更大元素 II
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	stack := make([]int, n+1)
	k := 0
	stack[k] = -1

	for i := 2*n - 1; i >= 0; i-- {
		ii := i % n
		for top := stack[k]; k != 0 && nums[ii] >= top; top = stack[k] {
			k--
		}
		ans[ii] = stack[k]
		k++
		stack[k] = nums[ii]
	}
	return ans
}

func testOne(in string, ans string) {
	nums := helper.ParseArray(in)
	result := nextGreaterElements(nums)
	helper.Assert(helper.DumpArray(result) == ans)
	helper.Log(in, "=>", ans)
}

func main() {
	testOne("[1,2,1]", "[2,-1,2]")
}
