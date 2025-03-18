package main

import (
	"leetcode/helper"
)

// 打开转盘锁
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	visited := make(map[string]bool)
	for _, s := range deadends {
		if s == "0000" {
			return -1
		}
		visited[s] = true
	}
	visited["0000"] = true
	queue := []string{"0000"}

	var step int
	for len(queue) > 0 {
		step++
		var tmp []string
		for _, s := range queue {
			nums := []byte(s)
			for _, newNum := range next(nums, visited) {
				if newNum == target {
					return step
				}
				visited[newNum] = true
				tmp = append(tmp, newNum)
			}
		}
		queue = tmp
	}
	return -1
}

func next(nums []byte, visited map[string]bool) []string {
	var newNums []string
	for i, c := range nums {
		// 加
		if c == '9' {
			nums[i] = '0'
		} else {
			nums[i] = c + 1
		}
		s := string(nums)
		if !visited[s] {
			newNums = append(newNums, s)
		}
		// 减
		if c == '0' {
			nums[i] = '9'
		} else {
			nums[i] = c - 1
		}
		s = string(nums)
		if !visited[s] {
			newNums = append(newNums, s)
		}
		nums[i] = c
	}
	return newNums
}

func testOne(in string, target string, ans int) {
	deadends := helper.ParseStrArray(in)
	helper.Assert(openLock(deadends, target) == ans)
}

func main() {
	testOne(`["0201","0101","0102","1212","2002"]`, "0202", 6)
	testOne(`["8888"]`, "0009", 1)
	testOne(`["8887","8889","8878","8898","8788","8988","7888","9888"]`, "8888", -1)
}
