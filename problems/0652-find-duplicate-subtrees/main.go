package main

import (
	"fmt"
	"leetcode/helper"
)

type TreeNode = helper.TreeNode

// 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	counter := make(map[string]int)
	var ans []*TreeNode
	var dfs func(*TreeNode) string

	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		serial := fmt.Sprintf("%d,%s,%s", node.Val, dfs(node.Left), dfs(node.Right))
		counter[serial]++
		if counter[serial] == 2 {
			ans = append(ans, node)
		}
		return serial
	}

	dfs(root)
	return ans
}

func testOne(in string) {
	t := helper.ParseTree(in)
	result := findDuplicateSubtrees(t)
	helper.Log(len(result), result[0], result[1])
}

func main() {
	testOne("[1,2,3,4,null,2,4,null,null,4]")
}
