package main

import (
	"leetcode/helper"
)

type Node = helper.Node

/*****************
        1
      /   \
     2     3
    / \   / \
   4   5 6   7
*****************/

// 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	for leftMost := root; leftMost.Left != nil; leftMost = leftMost.Left {
		for node := leftMost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	traverseConnect(root.Left, root.Right)
	return root
}

func traverseConnect(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	traverseConnect(left.Left, left.Right)
	traverseConnect(right.Left, right.Right)
	traverseConnect(left.Right, right.Left)
}

func main() {

}
