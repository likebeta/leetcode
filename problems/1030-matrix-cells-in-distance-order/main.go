package main

import (
	"leetcode/helper"
)

// 距离顺序排列矩阵单元格
var dir4 = [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func allCellsDistOrder(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 1, n*m)
	ans[0] = []int{r0, c0}
	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
	row, col := r0, c0
	for dist := 1; dist <= maxDist; dist++ {
		row--
		for i, dir := range dir4 {
			for i%2 == 0 && row != r0 || i%2 == 1 && col != c0 {
				if 0 <= row && row < n && 0 <= col && col < m {
					ans = append(ans, []int{row, col})
				}
				row += dir[0]
				col += dir[1]
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testOne(n, m int, r0, c0 int) {
	ans := allCellsDistOrder(n, m, r0, c0)
	helper.Log(helper.Dump(ans))
}

func main() {
	testOne(1, 2, 0, 0)
	testOne(2, 2, 0, 1)
	testOne(2, 3, 1, 2)
}
