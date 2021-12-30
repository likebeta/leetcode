package main

import (
	"leetcode/helper"
	"sort"
)

// 一手顺子
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	counter := make(map[int]int)
	for _, v := range hand {
		counter[v]++
	}
	for _, v := range hand {
		if counter[v] > 0 {
			for i := 0; i < groupSize; i++ {
				if counter[v+i] == 0 {
					return false
				}
				counter[v+i]--
			}
		}
	}
	return true
}

func testOne(in string, groupSize int, ans bool) {
	hand := helper.ParseArray(in)
	helper.Assert(isNStraightHand(hand, groupSize) == ans)
}

func main() {
	testOne("[1,2,3,6,2,3,4,7,8]", 3, true)
	testOne("[1,2,3,4,5]", 3, false)
	testOne("[1,2,3,4,5]", 4, false)
	testOne("[1,1,2,2,3,3]", 2, false)
	testOne("[1,1,2,2,3,3]", 3, true)
	testOne("[2,1]", 2, true)
}
