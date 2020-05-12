package main

import (
	"fmt"
	"leetcode/helper"
)

// 快乐数
func isHappy(n int) bool {
	slow, fast := n, getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func isHappy2(n int) bool {
	// cycle_members = {4, 16, 37, 58, 89, 145, 42, 20}
	for n != 1 && n != 4 {
		n = getNext(n)
	}
	return n == 1
}

func getNext(n int) int {
	var sum, unit int
	for n != 0 {
		unit, n = n%10, n/10
		sum += unit * unit
	}
	return sum
}

func main() {
	helper.Assert(isHappy(19) == true)
	helper.Assert(isHappy2(19) == true)
	for i := 1; i <= 365; i++ {
		t1, t2 := isHappy(i), isHappy2(i)
		helper.Assert(t1 == t2)
		if t1 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
