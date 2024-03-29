package main

import (
	"leetcode/helper"
)

// 岛屿的周长
type pair struct {
	x, y int
}

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return
}

func testOne(in string, ans int) {
	var matrix [][]int
	helper.Load([]byte(in), &matrix)
	helper.Assert(islandPerimeter(matrix) == ans)
}

func main() {
	testOne("[[0,1,0,0], [1,1,1,0], [0,1,0,0], [1,1,0,0]]", 16)
}
