package main

import (
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

func testOne(in string, ans string) {
	t := helper.NewTree(in)
	t = convertBST(t)
	helper.Log(t.Dump(), ans)
}

func main() {
	testOne("[5, 2, 13]", "[18,20,13]")
	testOne("[2,0,3,-4,1]", "[5,6,3,2,6]")

}
