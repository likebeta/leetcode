package main

import (
	"leetcode/helper"
)

// 滑动谜题
func slidingPuzzle(board [][]int) int {
	target := "123450"
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	i, sBorad := toStr(board)
	if sBorad == target {
		return 0
	}

	visited := make(map[string]bool)
	visited[sBorad] = true
	states := map[string]int{sBorad: i}

	var step int
	for len(states) > 0 {
		step++
		tmp := make(map[string]int)
		for s, i := range states {
			nums := []byte(s)
			for _, j := range neighbor[i] {
				nums[i], nums[j] = nums[j], nums[i]
				newS := string(nums)
				if newS == target {
					return step
				}
				if !visited[newS] {
					visited[newS] = true
					tmp[newS] = j
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		states = tmp
	}
	return -1
}

func toStr(board [][]int) (int, string) {
	var index0 int
	var sBorad []byte
	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == 0 {
				index0 = len(sBorad)
			}
			sBorad = append(sBorad, byte(board[i][j]+'0'))
		}
	}
	return index0, string(sBorad)
}

func testOne(in string, ans int) {
	matrix := helper.ParseIntMatrix(in)
	helper.Assert(slidingPuzzle(matrix) == ans)
}

func main() {
	testOne("[[1,2,3],[4,0,5]]", 1)
	testOne("[[1,2,3],[5,4,0]]", -1)
	testOne("[[4,1,2],[5,0,3]]", 5)
	testOne("[[1,2,3],[4,5,0]]", 0)
}
