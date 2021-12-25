package main

import (
	"leetcode/helper"
)

// 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 31; num > 0; i-- {
		ans += (num & 1) << i
		num >>= 1
	}
	return ans
}

func main() {
	helper.Assert(reverseBits(0b00000010100101000001111010011100) == 0b00111001011110000010100101000000)
	helper.Assert(reverseBits(0b11111111111111111111111111111101) == 0b10111111111111111111111111111111)
}
