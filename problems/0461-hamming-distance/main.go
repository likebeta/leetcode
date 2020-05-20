package main

import (
	"leetcode/helper"
)

// 汉明距离
func hammingDistance(x int, y int) int {
	num := x ^ y
	var ans int
	for num > 0 {
		ans++
		num &= num - 1
	}
	return ans
}

func main() {
	helper.Assert(hammingDistance(1, 4) == 2)
}
