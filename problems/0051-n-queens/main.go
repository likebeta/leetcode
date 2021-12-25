package main

import (
	"leetcode/helper"
)

// N 皇后
func solveNQueens(n int) [][]string {
	var ans [][]string
	if n < 1 {
		return ans
	}

	var (
		shuArr = make([]bool, n, n)
		pieArr = make([]bool, 2*n-1, 2*n-1)
		naArr = make([]bool, 2*n-1, 2*n-1)
		record = make([]int, n, n)
		process func(int)
	)

	process = func(row int) {
		if row == n {
			result := make([]string, n)
			for i := 0; i < n; i++ {
				one := make([]byte, n)
				for j := 0; j < n; j++ {
					one[j] = '.'
				}
				one[record[i]] = 'Q'
				result[i] = string(one)
			}
			ans = append(ans, result)
		}

		for col := 0; col < n; col++ {
			pie := row + col
			na := n - 1 - row + col
			if shuArr[col] || pieArr[pie] || naArr[na] {
				continue
			}
			record[row] = col
			shuArr[col] = true
			pieArr[pie] = true
			naArr[na] = true
			process(row+1)
			shuArr[col] = false
			pieArr[pie] = false
			naArr[na] = false
		}
	}
	process(0)
	return ans
}

func main() {
	helper.Log(solveNQueens(4))
}
