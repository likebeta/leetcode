package main

import (
	"encoding/json"
	"leetcode/helper"
)

// 二叉树的序列化与反序列化
type TreeNode = helper.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	i, last, queue := 0, 0, []*TreeNode{root}
	for i < len(queue) {
		if queue[i] != nil {
			last = i
			queue = append(queue, queue[i].Left, queue[i].Right)
		}
		i++
	}
	arr := make([]*int, last+1, last+1)
	for i := range arr {
		if queue[i] != nil {
			arr[i] = &queue[i].Val
		}
	}
	data, _ := json.Marshal(arr)
	return string(data)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var arr []*int
	if err := json.Unmarshal([]byte(data), &arr); err != nil {
		return nil
	}

	var root *TreeNode
	cnt := len(arr)
	if cnt > 0 {
		root = &TreeNode{Val: *arr[0]}
		arr = arr[1:]
		list := []*TreeNode{root}
		for len(list) > 0 && len(arr) > 0 {
			curr := list[0]
			list = list[1:]
			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Left = &TreeNode{Val: *arr[0]}
					list = append(list, curr.Left)
				}
				arr = arr[1:]
			}

			if len(arr) > 0 {
				if arr[0] != nil {
					curr.Right = &TreeNode{Val: *arr[0]}
					list = append(list, curr.Right)
				}
				arr = arr[1:]
			}
		}
	}
	return root
}

func testOne(s string) {
	root := helper.NewTree(s)
	obj := Constructor()
	data := obj.serialize(root)
	helper.Log(data, s)
	ans := obj.deserialize(data)
	helper.Log(ans.Dump(), s)
}

func main() {
	testOne("[]")
	//testOne("[1,2,3,null,null,4,5]")
}
