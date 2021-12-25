package main

import (
	"leetcode/helper"
)

// 最低票价
func mincostTickets(days []int, costs []int) int {
	nDays := days[len(days)-1]
	dp := make([]int, nDays+31, nDays+31)
	dMap := make(map[int]bool)
	for i := range days {
		dMap[days[i]] = true
	}
	for i := nDays; i > 0; i-- {
		if dMap[i] {
			dp[i] = min(dp[i+1]+costs[0], dp[i+7]+costs[1], dp[i+30]+costs[2])
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[1]
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

func mincostTickets2(days []int, costs []int) int {
	VALIDITY := []int{1, 7, 30}
	nDays := len(days)
	dp := make([]int, nDays+30, nDays+30)
	dp[0] = costs[0] * days[nDays-1]
	for i := nDays - 1; i >= 0; i-- {
		dp[i] = dp[0]
		for j := range VALIDITY {
			k, end := j+1, days[i]+VALIDITY[j]
			for k < nDays && end > days[k] {
				k++
			}
			if costs[j]+dp[k] < dp[i] {
				dp[i] = costs[j] + dp[k]
			}
		}
	}
	return dp[0]
}

func main() {
	var days, costs []int
	days = []int{1, 4, 6, 7, 8, 20}
	costs = []int{2, 7, 15}
	helper.Assert(mincostTickets(days, costs) == 11 && mincostTickets2(days, costs) == 11)
	days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs = []int{2, 7, 15}
	helper.Assert(mincostTickets(days, costs) == 17 && mincostTickets2(days, costs) == 17)
	days = []int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28}
	costs = []int{3, 13, 45}
	helper.Assert(mincostTickets(days, costs) == 44 && mincostTickets2(days, costs) == 44)
}
