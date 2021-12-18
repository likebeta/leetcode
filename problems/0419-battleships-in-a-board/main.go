package main

import (
	"leetcode/helper"
)

// 甲板上的战舰
func countBattleships(board [][]byte) int {
	var ans int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 'X' || (i > 0 && board[i-1][j] == 'X') || (j > 0 && board[i][j-1] == 'X') {
				continue
			}
			ans++
		}
	}
	return ans
}

func testOne(in string, ans int) {
	var grid [][]string
	helper.Load([]byte(in), &grid)
	matrix := make([][]byte, len(grid))
	for row := 0; row < len(grid); row++ {
		matrix[row] = make([]byte, len(grid[0]))
		for col := 0; col < len(grid[0]); col++ {
			matrix[row][col] = grid[row][col][0]
		}
	}
	result := countBattleships(matrix)
	helper.Assert(result == ans)
}

func main() {
	testOne(`[["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]`, 2)
	testOne(`[["."]]`, 0)
}
