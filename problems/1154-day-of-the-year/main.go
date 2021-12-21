package main

import (
	"leetcode/helper"
	"strconv"
)

// 一年中的第几天
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		monthDays[1]++
	}
	ans := day
	for i := 0; i < month-1; i++ {
		ans += monthDays[i]
	}
	return ans
}

func testOne(in string, ans int) {
	days := dayOfYear(in)
	helper.Assert(days == ans)
	helper.Log(in, "=>", days)
}

func main() {
	testOne("2019-01-09", 9)
	testOne("2019-02-10", 41)
	testOne("2003-03-01", 60)
	testOne("2004-03-01", 61)
}
