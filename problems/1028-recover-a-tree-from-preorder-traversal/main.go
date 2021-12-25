package main

import (
	"leetcode/helper"
	"strconv"
)

type TreeNode = helper.TreeNode

// 从先序遍历还原二叉树
func recoverFromPreorder(S string) *TreeNode {
	if len(S) == 0 {
		return nil
	}
	index := 0
	for _, v := range S {
		if v != '-' {
			index++
		} else {
			break
		}
	}
	sum := 0
	i, _ := strconv.Atoi(S[:index])
	root := &TreeNode{Val: i}
	mark := &TreeNode{Val: 0, Left: root}
	parent := mark

	substring := S[index:]
	for i := 0; i < len(substring); i++ {
		if substring[i] == '-' {
			//找当前结点的父结点, 每发现 - 从root向下搜索一个结点, 因为right最后被设置, 优先使用right
			if parent.Right != nil {
				parent = parent.Right
			} else {
				parent = parent.Left
			}
			sum++
			continue
		} else {
			tail := i + 1
			for ; tail < len(substring); tail++ {
				if substring[tail] == '-' {
					break
				}
			}
			v, _ := strconv.Atoi(substring[i:tail])
			node := &TreeNode{Val: v}
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
			i = tail - 1
			parent = mark
		}
	}
	return root
}

func main() {
	helper.Assert(recoverFromPreorder("1-2--3--4-5--6--7").Dump() == "[1,2,5,3,4,6,7]")
	helper.Assert(recoverFromPreorder("1-2--3---4-5--6---7").Dump() == "[1,2,5,3,null,6,null,4,null,7]")
	helper.Assert(recoverFromPreorder("1-401--349---90--88").Dump() == "[1,401,null,349,88,90]")
}
