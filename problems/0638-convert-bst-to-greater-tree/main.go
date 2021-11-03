package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += sum
		sum = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func testOne(in string, ans string)  {
	t := helper.ParseTree(in)
	t = convertBST(t)
	helper.Log(in, "=>", t.Dump() == ans)
}

func main() {
	testOne("[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]", "[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]")
	testOne("[0,null,1]", "[1,null,1]")
	testOne("[1,0,2]", "[3,3,2]")
	testOne("[3,2,4,1]", "[7,9,4,10]")
}
