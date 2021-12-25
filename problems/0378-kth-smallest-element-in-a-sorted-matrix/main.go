package main

import (
	"leetcode/helper"
)

// 有序矩阵中第K小的元素
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}

func main() {
	var matrix [][]int
	matrix = [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	helper.Assert(kthSmallest(matrix, 8) == 13)
}
