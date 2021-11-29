package main

import (
	"leetcode/helper"
)

// 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if k == 1 || n < k {
		return nums
	}
	ans := make([]int, n-k+1)
	queue := make([]int, 0)
	for i, v := range nums {
		l := len(queue)
		for l > 0 && nums[queue[l-1]] < v {
			l--
			queue = queue[:l]
		}
		queue = append(queue, i)
		if i-k+1 >= 0 {
			ans[i-k+1] = nums[queue[0]]
			if queue[0] == i-k+1 {
				queue = queue[1:]
			}
		}
	}
	return ans
}

func testOne(in string, k int, ans string) {
	nums := helper.ParseArray(in)
	result := maxSlidingWindow(nums, k)
	helper.Assert(helper.DumpArray(result) == ans)
	helper.Log(in, k, "=>", ans)
}

func main() {
	testOne("[1,3,-1,-3,5,3,6,7]", 3, "[3,3,5,5,6,7]")
	testOne("[1]", 1, "[1]")
	testOne("[1,-1]", 1, "[1,-1]")
	testOne("[9,11]", 2, "[11]")
	testOne("[4,-2]", 2, "[4]")
}
