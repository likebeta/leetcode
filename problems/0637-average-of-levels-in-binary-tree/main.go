package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	var (
		ans   []float64
		queue []*TreeNode
	)

	if root != nil {
		queue = append(queue, root)
		for len(queue) > 0 {
			n := len(queue)
			total := float64(0)
			for i := 0; i < n; i++ {
				total += float64(queue[i].Val)
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			ans = append(ans, total/float64(n))
			queue = queue[n:]
		}
	}

	return ans
}

func main() {
	t := helper.NewTree("[3, 9, 20, null, null, 15, 7]")
	t.Print(8)
	helper.Log(averageOfLevels(t))  // [3, 14.5, 11]
}
