package main

import (
	"leetcode/helper"
	"sort"
)

// 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	resultV := nums[length-1] + nums[length-2] + nums[length-3]
	if resultV == target {
		return target
	}
	absDiff := resultV - target
	if absDiff < 0 {
		absDiff = -absDiff
	}
	for i := 0; i < length-2; {
		front := i + 1
		end := length - 1
		for front < end {
			v := nums[i] + nums[front] + nums[end]
			if v == target {
				return target
			}
			d := v - target
			if d > 0 {
				for end--; front < end && nums[end] == nums[end+1]; end-- {
				}
				if d < absDiff {
					resultV = v
					absDiff = d
				}
			} else {
				for front++; front < end && nums[front] == nums[front-1]; front++ {
				}
				if -d < absDiff {
					resultV = v
					absDiff = -d
				}
			}
		}
		for i++; i < length-2 && nums[i] == nums[i-1]; i++ {
		}
	}
	return resultV
}

func main() {
	helper.Assert(threeSumClosest([]int{-1, 2, 1, -4}, 1) == 2)
}
