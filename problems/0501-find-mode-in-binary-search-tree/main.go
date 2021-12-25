package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 二叉搜索树中的众数
func findMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}

func main() {
	t := helper.NewTree("[1,null,2,2]")
	helper.Log(findMode(t), "[2]")
}
