package main

import (
	"encoding/json"
	"leetcode/helper"
	"log"
)

type Node = helper.Node

// 填充每个节点的下一个右侧节点指针 II
func connect(root *Node) *Node {
	for start := root; start != nil; {
		var first, prev *Node
		for curr := start; curr != nil; curr = curr.Next {
			if curr.Left != nil {
				if prev == nil {
					first, prev = curr.Left, curr.Left
				} else {
					prev.Next, prev = curr.Left, curr.Left
				}
			}
			if curr.Right != nil {
				if prev == nil {
					first, prev = curr.Right, curr.Right
				} else {
					prev.Next, prev = curr.Right, curr.Right

				}
			}
		}
		start = first
	}
	return root
}

func parse(bfs string) *Node {
	var arr []*int
	if err := json.Unmarshal([]byte(bfs), &arr); err != nil {
		log.Panic("parse failed:", err)
	}

	var root *Node
	cnt := len(arr)
	if cnt > 0 {
		root = &Node{Val: *arr[0]}
		arr = arr[1:]
		list := []*Node{root}
		for len(list) > 0 && len(arr) > 0 {
			curr := list[0]
			list = list[1:]
			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Left = &Node{Val: *arr[0]}
					list = append(list, curr.Left)
				}
				arr = arr[1:]
			}

			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Right = &Node{Val: *arr[0]}
					list = append(list, curr.Right)
				}
				arr = arr[1:]
			}
		}
	}
	return root
}

func main() {
	node := parse("[1,2,3,4,5,null,7]")
	node = connect(node)
}
