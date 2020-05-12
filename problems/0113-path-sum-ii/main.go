package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 路径总和 II
func pathSum(root *TreeNode, sum int) [][]int {
	return getPath(root, sum, nil, nil)
}

func getPath(root *TreeNode, sum int, result [][]int, path []int) [][]int {
	if root == nil {
		return result
	}

	sum -= root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			tmp := make([]int, len(path), len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}
		return result
	}
	result = getPath(root.Left, sum, result, path)
	result = getPath(root.Right, sum, result, path)
	return result
}

func main() {
	var node *TreeNode
	node = helper.NewTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]")
	helper.PrintTree(node, 8)
	helper.Log(pathSum(node, 22)) // [[5 4 11 2] [5 8 4 5]]
}
