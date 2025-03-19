package main

import (
	"leetcode/helper"
	"slices"
)

// 下降路径最小和
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n)
	for row := range matrix {
		tmp := slices.Clone(dp)
		for col, x := range matrix[row] {
			// 状态转移：min(上一行上方，左上方，右上方) + x
			if col > 0 {
				tmp[col] = min(tmp[col], dp[col-1])
			}
			if col+1 < n {
				tmp[col] = min(tmp[col], dp[col+1])
			}
			tmp[col] += x
		}
		dp = tmp
	}
	return slices.Min(dp)
}

func testOne(in string, ans int) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(minFallingPathSum(matrix) == ans)
}

func main() {
	testOne("[[2,1,3],[6,5,4],[7,8,9]]", 13)
	testOne("[[-19,57],[-40,-5]]", -59)
}
