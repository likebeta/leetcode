package main

import (
	"leetcode/helper"
	"sort"
)

// 供暖器
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	var ans int
	j := 0
	for i := range houses {
		dis := abs(houses[i] - heaters[j])
		for j+1 < len(heaters) && dis >= abs(houses[i]-heaters[j+1]) {
			dis = abs(houses[i] - heaters[j+1])
			j++
		}
		if dis > ans {
			ans = dis
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func testOne(in1, in2 string, ans int) {
	houses := helper.ParseArray(in1)
	heaters := helper.ParseArray(in2)
	result := findRadius(houses, heaters)
	helper.LogF("[%s, %s] => %d", in1, in2, result)
	helper.Assert(result == ans)
}

func main() {
	testOne("[1,2,3]", "[2]", 1)
	testOne("[1,2,3,4]", "[1,4]", 1)
	testOne("[1,5]", "[2]", 3)
}
