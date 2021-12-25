package main

import (
	"fmt"
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root != nil {
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			levelNodes := make([]int, len(queue))
			for i := 0; i < len(levelNodes); i++ {
				fmt.Println(i, queue)
				levelNodes[i] = queue[i].Val
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}

				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[len(levelNodes):]
			ans = append(ans, levelNodes)
		}

		for i := 0; i < len(ans)/2; i++ {
			ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
		}
	}
	return ans
}

func main() {
	t := helper.NewTree("[3,9,20,null,null,15,7]")
	helper.Log(levelOrderBottom(t))
}
