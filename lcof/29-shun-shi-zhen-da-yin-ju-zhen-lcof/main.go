package main

import (
	"leetcode/helper"
)

// 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col := len(matrix), len(matrix[0])
	total := row * col
	ans := make([]int, total, total)
	l, r, t, b := 0, col-1, 0, row-1
	for i := 0; i < total; {
		for j := l; j <= r; j++ {
			i, ans[i] = i+1, matrix[t][j]
		}
		for j := t + 1; j <= b; j++ {
			i, ans[i] = i+1, matrix[j][r]
		}
		if l < r && t < b {
			for j := r - 1; j > l; j-- {
				i, ans[i] = i+1, matrix[b][j]
			}
			for j := b; j > t; j-- {
				i, ans[i] = i+1, matrix[j][l]
			}
		}
		l, r, t, b = l+1, r-1, t+1, b-1
	}
	return ans
}

func main() {
	var m [][]int
	m = [][]int{
		{3},
		{2},
	}
	helper.Log(spiralOrder(m), []int{3, 2})
	m = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	helper.Log(spiralOrder(m), []int{1, 2, 4, 6, 5, 3})
	m = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	helper.Log(spiralOrder(m), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	m = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	helper.Log(spiralOrder(m), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}
