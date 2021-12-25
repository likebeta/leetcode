package main

import (
	"leetcode/helper"
)

// 单词搜索
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	var (
		h     = len(board)
		w     = len(board[0])
		vis   = make([][]bool, h)
		check func(int, int, int) bool
	)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					vis[i][j] = false
					return true
				}
			}
		}
		vis[i][j] = false
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	helper.Assert(exist(board, "ABCCED") == true)
	helper.Assert(exist(board, "SEE") == true)
	helper.Assert(exist(board, "ABCB") == false)
}
