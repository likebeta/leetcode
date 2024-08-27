package main

import (
	"leetcode/helper"
)

// NumMatrix 二维区域和检索 - 矩阵不可变
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	sums := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		sums[i] = make([]int, col+1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row1][col2+1] - this.sums[row2+1][col1] + this.sums[row1][col1]
}

func main() {
	matrix := helper.ParseMatrix[int]("[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]")
	numArray := Constructor(matrix)
	helper.Assert(numArray.SumRegion(2, 1, 4, 3) == 8)
	helper.Assert(numArray.SumRegion(1, 1, 2, 2) == 11)
	helper.Assert(numArray.SumRegion(1, 2, 2, 4) == 12)
}
