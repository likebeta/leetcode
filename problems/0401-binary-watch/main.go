package main

import (
	"fmt"
	"leetcode/helper"
)

// 二进制手表
func readBinaryWatch(num int) []string {
	var ans []string
	for hour := 0; hour < 12; hour++ {
		for minute := 0; minute < 60; minute++ {
			if count1(hour)+count1(minute) == num {
				ans = append(ans, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}
	}
	return ans
}

func count1(n int) int {
	var count int
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func readBinaryWatch2(num int) []string {
	var (
		ans       []string
		backtrace func(num, i, mask int)
	)

	backtrace = func(num, i, mask int) {
		if i > 10 {
			return
		}
		if num == 0 {
			hour, minute := mask>>6, mask&0b111111
			if hour < 12 && minute < 60 {
				ans = append(ans, fmt.Sprintf("%d:%02d", hour, minute))
			}
			return
		}

		for k := i; k < 10; k++ {
			if (2<<k)&mask == 0 {
				backtrace(num-1, k+1, mask|(1<<k))
			}
		}
	}
	backtrace(num, 0, 0)
	return ans
}

func main() {
	helper.Log(readBinaryWatch(1))
	helper.Log(readBinaryWatch2(1))
	helper.Log(readBinaryWatch(8))
	helper.Log(readBinaryWatch2(8))
}
