package main

import (
	"leetcode/helper"
)

// 位1的个数
func hammingWeight(num uint32) int {
	var ans int
	for num > 0 {
		ans++
		num &= num - 1
	}
	return ans
}

func main() {
	helper.Assert(hammingWeight(0b00000000000000000000000000001011) == 3)
	helper.Assert(hammingWeight(0b00000000000000000000000010000000) == 1)
	helper.Assert(hammingWeight(0b11111111111111111111111111111101) == 31)
}
