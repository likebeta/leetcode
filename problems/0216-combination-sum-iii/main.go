package main

import (
	"leetcode/helper"
)

// 组合总和 III
func combinationSum3(k int, n int) [][]int {
	var (
		ans    [][]int
		dfs    func(int, int, int)
		record = make([]int, k)
	)
	dfs = func(i int, j int, n int) {
		if n == 0 && i == k {
			dest := make([]int, k)
			copy(dest, record)
			ans = append(ans, dest)
			return
		}
		if n == 0 || i == k || j > 9 || j > n || k-i > 10-j {
			return
		}
		record[i] = j
		dfs(i+1, j+1, n-j)
		dfs(i, j+1, n)
	}
	dfs(0, 1, n)
	return ans
}

func main() {
	helper.Log(combinationSum3(3, 7)) // [[1,2,4]]
	helper.Log(combinationSum3(3, 9)) // [[1,2,6], [1,3,5], [2,3,4]]
}
