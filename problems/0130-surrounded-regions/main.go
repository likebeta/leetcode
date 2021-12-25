package main

import (
	"fmt"
)

func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	var (
		M   = len(board)
		N   = len(board[0])
		dfs func(int, int)
	)

	dfs = func(m, n int) {
		if m < 0 || m >= M || n < 0 || n >= N || board[m][n] != 'O' {
			return
		}
		board[m][n] = 'K'
		dfs(m+1, n)
		dfs(m-1, n)
		dfs(m, n+1)
		dfs(m, n-1)
	}

	for n := 0; n < N; n++ {
		dfs(0, n)
		dfs(M-1, n)
	}

	for m := M - 2; m > 0; m-- {
		dfs(m, 0)
		dfs(m, N-1)
	}

	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			if board[m][n] == 'K' {
				board[m][n] = 'O'
			} else if board[m][n] == 'O' {
				board[m][n] = 'X'
			}
		}
	}
}

func testOne(board [][]byte) {
	var cache []string
	for i := range board {
		cache = append(cache, string(board[i]))
	}
	solve(board)
	for i := range board {
		fmt.Print(cache[i], "--->", string(board[i]), "\n")
	}
}

func main() {
	testOne([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	})
}
