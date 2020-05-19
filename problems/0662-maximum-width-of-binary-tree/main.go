package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

type One struct {
	Node *TreeNode
	Pos  int
}

// 二叉树最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	queue := []*One{{root, 0}}
	for len(queue) > 0 {
		if queue[len(queue)-1].Pos-queue[0].Pos > ans {
			ans = queue[len(queue)-1].Pos - queue[0].Pos
		}
		var lastQueue []*One
		for i := 0; i < len(queue); i++ {
			if queue[i].Node.Left != nil {
				lastQueue = append(lastQueue, &One{queue[i].Node.Left, queue[i].Pos * 2})
			}
			if queue[i].Node.Right != nil {
				lastQueue = append(lastQueue, &One{queue[i].Node.Right, queue[i].Pos*2 + 1})
			}
		}
		queue = lastQueue
	}

	return ans + 1
}

func main() {
	var t *TreeNode
	t = helper.NewTree("[1,3,2,5,3,null,9]")
	helper.PrintTree(t, 8)
	helper.Assert(widthOfBinaryTree(t) == 4)

	t = helper.NewTree("[1,3,null,5,3]")
	helper.PrintTree(t, 8)
	helper.Assert(widthOfBinaryTree(t) == 2)

	t = helper.NewTree("[1,3,2,5,null,null,9,6,null,null,7]")
	helper.PrintTree(t, 8)
	helper.Assert(widthOfBinaryTree(t) == 8)

	t = helper.NewTree("[1,3,null,null,3]")
	helper.PrintTree(t, 8)
	helper.Assert(widthOfBinaryTree(t) == 1)
}
