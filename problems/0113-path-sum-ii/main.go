package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 路径总和 II
func pathSum(root *TreeNode, sum int) [][]int {
	var (
		ans     [][]int
		getPath func(*TreeNode, int, []int)
	)

	getPath = func(node *TreeNode, sum int, record []int) {
		if node != nil {
			sum -= node.Val
			record = append(record, node.Val)
			if node.Left == nil && node.Right == nil {
				if sum == 0 {
					tmp := make([]int, len(record))
					copy(tmp, record)
					ans = append(ans, tmp)
				}
			} else {
				getPath(node.Left, sum, record)
				getPath(node.Right, sum, record)
			}
		}
	}

	getPath(root, sum, nil)
	return ans
}

func main() {
	t := helper.NewTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]")
	helper.Log(pathSum(t, 22)) // [[5 4 11 2] [5 8 4 5]]
}
