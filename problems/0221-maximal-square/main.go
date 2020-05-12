package main

import (
	"leetcode/helper"
)

// 最大正方形
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxSide int
	row, col := len(matrix), len(matrix[0])
	var preLeftTop int
	line := make([]int, col, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			preLeftTop, line[j] = line[j], preLeftTop
			if matrix[i][j] == '0' {
				line[j] = 0
				continue
			}
			if i == 0 || j == 0 {
				line[j] = 1
			} else {
				line[j] = min(line[j], line[j-1], preLeftTop) + 1
			}
			if line[j] > maxSide {
				maxSide = line[j]
			}
		}

	}
	return maxSide * maxSide
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

func main() {
	var matrix [][]byte
	matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	helper.Assert(maximalSquare(matrix) == 4)

	matrix = [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	helper.Assert(maximalSquare(matrix) == 9)
}
