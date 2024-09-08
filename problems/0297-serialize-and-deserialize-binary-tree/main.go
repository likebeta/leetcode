package main

import (
	"encoding/json"
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// Codec 二叉树的序列化与反序列化
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
	nodes, last := []*TreeNode{root}, 0
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			last = i
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		}
	}
	arr := make([]*int, last+1)
	for i := range arr {
		if nodes[i] != nil {
			arr[i] = &nodes[i].Val
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

	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	parents := []*TreeNode{root}

	for j, i := 0, 1; j < len(parents) && i < len(arr); j, i = j+1, i+2 {
		if arr[i] != nil {
			parents[j].Left = &TreeNode{Val: *arr[i]}
			parents = append(parents, parents[j].Left)
		}
		if i+1 < len(arr) && arr[i+1] != nil {
			parents[j].Right = &TreeNode{Val: *arr[i+1]}
			parents = append(parents, parents[j].Right)
		}
	}
	return root
}

func testOne(s string) {
	root := helper.NewTree(s)
	obj := Constructor()
	encoded := obj.serialize(root)
	helper.Assert(encoded == s)
	decoded := obj.deserialize(encoded)
	helper.Assert(decoded.Dump() == s)
}

func main() {
	testOne("[]")
	testOne("[1,2,3,null,null,4,5]")
	testOne("[1,2]")
}
