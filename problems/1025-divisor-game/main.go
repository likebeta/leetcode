package main

import (
	"leetcode/helper"
)

// 除数博弈
func divisorGame(N int) bool {
	return N % 2 == 0
}

func main() {
	helper.Assert(divisorGame(1) == false)
	helper.Assert(divisorGame(2) == true)
	helper.Assert(divisorGame(3) == false)
	helper.Assert(divisorGame(4) == true)
}
