package main

import (
	"leetcode/helper"
)

// x 的平方根
func mySqrt(x int) int {
	i := 1
	for i*i <= x {
		i++
	}
	return i - 1
}

func mySqrt2(x int) int {
	var left, mid int
	right := x
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func main() {
	for i := 0; i < 10000; i++ {
		helper.Log(i, "==>", mySqrt(i), mySqrt2(i))
		helper.Assert(mySqrt(i) == mySqrt2(i))
	}
}
