package main

import (
	"leetcode/helper"
)

// 机器人能否返回原点
func judgeCircle(moves string) bool {
	var x, y int
	for i := range moves {
		switch moves[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}

func main() {
	helper.Assert(judgeCircle("UD") == true)
	helper.Assert(judgeCircle("LL") == false)
}
