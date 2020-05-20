package main

import (
	"leetcode/helper"
)

// 将数字变成 0 的操作次数
func numberOfSteps(num int) int {
	var count int
	for num != 0 {
		if num&1 == 1 {
			num -= 1
		} else {
			num >>= 1
		}
		count++
	}
	return count
}

func numberOfSteps2(num int) int {
	var count int
	for num > 1 {
		if num&1 == 1 {
			count += 2
		} else {
			count++
		}
		num >>= 1
	}
	return count + num
}

func main() {
	helper.Log(numberOfSteps(14) == 6, numberOfSteps2(14) == 6)
	helper.Log(numberOfSteps(8) == 4, numberOfSteps2(8) == 4)
	helper.Log(numberOfSteps(123) == 12, numberOfSteps2(123) == 12)
}
