package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root != nil {
		i0, i1 := 0, 1
		var arr [2][]*TreeNode
		arr[0] = append(arr[0], root)
		for len(arr[i0]) > 0 {
			var oneLevel []int
			for i := range arr[i0] {
				oneLevel = append(oneLevel, arr[i0][i].Val)
				if arr[i0][i].Left != nil {
					arr[i1] = append(arr[i1], arr[i0][i].Left)
				}
				if arr[i0][i].Right != nil {
					arr[i1] = append(arr[i1], arr[i0][i].Right)
				}
			}
			ans = append(ans, oneLevel)
			arr[i0] = nil
			i0, i1 = i1, i0
		}
	}
	return ans
}

func main() {
	t := helper.NewTree("[3,9,20,null,null,15,7]")
	helper.PrintTree(t, 6)
	helper.Log(levelOrder(t))
}
