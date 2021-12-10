package main

import (
	"leetcode/helper"
)

// 旋转图像

/*
分块，遍历左上角区域
temp                    =matrix[row][col]
matrix[row][col]        =matrix[n−col−1][row]
matrix[n−col−1][row]    =matrix[n−row−1][n−col−1]
matrix[n−row−1][n−col−1]=matrix[col][n−row−1]
matrix[col][n−row−1]    =temp
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

/*
分层，遍历上部倒梯形
*/
func rotate2(matrix [][]int) {
	n := len(matrix)
	start, level := 0, n-1
	for start < level {
		for j := start; j < level; j++ {
			temp := matrix[start][j]
			matrix[start][j] = matrix[n-1-j][start]
			matrix[n-1-j][start] = matrix[n-1-start][n-1-j]
			matrix[n-1-start][n-1-j] = matrix[j][n-1-start]
			matrix[j][n-1-start] = temp
		}
		start++
		level--
	}
}

func testOne(in string, ans string) {
	{
		var matrix [][]int
		helper.Load([]byte(in), &matrix)
		rotate(matrix)
		out := helper.Dump(&matrix)
		helper.Assert(out == ans)
	}
	{
		var matrix [][]int
		helper.Load([]byte(in), &matrix)
		rotate2(matrix)
		out := helper.Dump(&matrix)
		helper.Assert(out == ans)
	}
}
func main() {
	testOne("[[1,2,3],[4,5,6],[7,8,9]]", "[[7,4,1],[8,5,2],[9,6,3]]")
	testOne("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]", "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]")
	testOne("[[1]]", "[[1]]")
	testOne("[[1,2],[3,4]]", "[[3,1],[4,2]]")
}
