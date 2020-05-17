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

func findDuplicate2(nums []int) int {
	// 二分查找，限制1和n必须存在时方法有效
	left, right := 1, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		var count int
		for i := range nums {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	var arr []int
	arr = []int{1, 3, 4, 2, 2}
	helper.Assert(findDuplicate(arr) == 2)
	helper.Assert(findDuplicate2(arr) == 2)
	arr = []int{3, 1, 3, 4, 2}
	helper.Assert(findDuplicate(arr) == 3)
	helper.Assert(findDuplicate2(arr) == 3)
}
