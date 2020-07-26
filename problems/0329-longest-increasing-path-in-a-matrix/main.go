package main

import "leetcode/helper"

// 矩阵中的最长递增路径
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for row := 0; row < rows; row++ {
		memo[row] = make([]int, cols)
	}
	var (
		ans  int
		dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		dfs  func(int, int) int
	)

	dfs = func(row int, col int) int {
		if memo[row][col] != 0 {
			return memo[row][col]
		}
		memo[row][col]++
		for _, dir := range dirs {
			r, c := row+dir[0], col+dir[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && matrix[r][c] > matrix[row][col] {
				memo[row][col] = max(memo[row][col], dfs(r, c)+1)
			}
		}
		return memo[row][col]
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans = max(ans, dfs(row, col))
		}
	}
	return ans
}

func longestIncreasingPath2(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(matrix), len(matrix[0])
	inDegree := make([][]int, rows)
	for i := 0; i < rows; i++ {
		inDegree[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range dirs {
				r, c := i+dir[0], j+dir[1]
				if r >= 0 && r < rows && c >= 0 && c < cols && matrix[r][c] < matrix[i][j] {
					inDegree[i][j]++
				}
			}
		}
	}

	var queue [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inDegree[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	ans := 0
	for len(queue) != 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]
			row, column := cell[0], cell[1]
			for _, dir := range dirs {
				r, c := row+dir[0], column+dir[1]
				if r >= 0 && r < rows && c >= 0 && c < cols && matrix[r][c] > matrix[row][column] {
					inDegree[r][c]--
					if inDegree[r][c] == 0 {
						queue = append(queue, []int{r, c})
					}
				}
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func testOne(matrix [][]int, ans int) {
	helper.Assert(longestIncreasingPath(matrix) == ans)
	helper.Assert(longestIncreasingPath2(matrix) == ans)

}

func main() {
	testOne([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}, 4)
	testOne([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}, 4)
}
