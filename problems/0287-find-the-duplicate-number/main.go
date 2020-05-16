package main

import (
	"leetcode/helper"
)

// 寻找重复数
func findDuplicate(nums []int) int {
	// 把数组下标和对象的值都理解为指针，会有环
	var slow, fast int
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	helper.Assert(findDuplicate([]int{1, 3, 4, 2, 2}) == 2)
	helper.Assert(findDuplicate([]int{3, 1, 3, 4, 2}) == 3)
}
