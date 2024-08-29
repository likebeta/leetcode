package main

import (
	"leetcode/helper"
)

// 航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		i, j, seats := booking[0], booking[1], booking[2]
		ans[i-1] += seats
		if j < n {
			ans[j] -= seats
		}
	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

func testOne(length int, updates string, ans string) {
	matrix := helper.ParseMatrix[int](updates)
	ret := corpFlightBookings(matrix, length)
	helper.Assert(helper.DumpArray(ret) == ans)
}

func main() {
	testOne(5, "[[1,2,10],[2,3,20],[2,5,25]]", "[10,55,45,25,25]")
	testOne(2, "[[1,2,10],[2,2,15]]", "[10,25]")
}
