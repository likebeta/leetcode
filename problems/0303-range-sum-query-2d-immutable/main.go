package main

import (
	"leetcode/helper"
)

// NumArray 区域和检索 - 数组不可变
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = nums[i] + sums[i]
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

func main() {
	arr := helper.ParseArray("[-2,0,3,-5,2,-1]")
	numArray := Constructor(arr)
	helper.Assert(numArray.SumRange(0, 2) == 1)
	helper.Assert(numArray.SumRange(2, 5) == -1)
	helper.Assert(numArray.SumRange(0, 5) == -3)
}
