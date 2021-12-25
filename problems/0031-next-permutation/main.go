package main

import (
	"leetcode/helper"
)

// 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i >= 0 {
		j := n - 1
		for ; j >= 0 && nums[i] >= nums[j]; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func testOne(in string, ans string) {
	nums := helper.ParseArray(in)
	nextPermutation(nums)
	helper.Assert(helper.DumpArray(nums) == ans)
}

func main() {
	testOne("[1,2,3]", "[1,3,2]")
	testOne("[1,3,2]", "[2,1,3]")
	testOne("[3,2,1]", "[1,2,3]")
}
