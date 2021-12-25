package main

import "leetcode/helper"

type TreeNode = helper.TreeNode

// 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	var (
		max     = ^(int(^uint(0) >> 1))
		process func(*TreeNode) int
	)
	process = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := process(node.Left)
		if leftGain < 0 {
			leftGain = 0
		}
		rightGain := process(node.Right)
		if rightGain < 0 {
			rightGain = 0
		}
		currGain := node.Val + leftGain + rightGain
		if currGain > max {
			max = currGain
		}
		if leftGain > rightGain {
			return node.Val + leftGain
		} else {
			return node.Val + rightGain
		}
	}
	process(root)
	return max
}

func testOne(bfs string, ans int) {
	t := helper.NewTree(bfs)
	helper.Assert(maxPathSum(t) == ans)
}

func main() {
	testOne("[1,2,3]", 6)
	testOne("[-10,9,20,null,null,15,7]", 42)
}
