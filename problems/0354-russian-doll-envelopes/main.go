package main

import (
	"leetcode/helper"
	"slices"
	"sort"
)

// 俄罗斯套娃信封问题 - O(n²)
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	// 按宽度升序排序，当宽度相同时按高度降序排序
	// 这样可以确保在宽度相同的情况下，只能选择一个信封（因为高度降序后不可能嵌套）
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1] // 宽度相同时，高度降序
		}
		return envelopes[i][0] < envelopes[j][0] // 宽度升序
	})

	// 问题转化为求高度数组的最长递增子序列
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		heights[i] = envelopes[i][1]
	}

	// 应用最长递增子序列算法
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // 最小长度为1
		for j := 0; j < i; j++ {
			if heights[i] > heights[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

// 优化版本 - 使用二分查找 - O(n log n)
func maxEnvelopes2(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	// 按宽度升序排序，当宽度相同时按高度降序排序
	// 这样可以确保在宽度相同的情况下，只能选择一个信封（因为高度降序后不可能嵌套）
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1] // 宽度相同时，高度降序
		}
		return envelopes[i][0] < envelopes[j][0] // 宽度升序
	})

	// 问题转化为求高度数组的最长递增子序列
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		heights[i] = envelopes[i][1]
	}

	// 使用二分查找优化的最长递增子序列算法 - O(n log n)
	// tails[i] 表示长度为i+1的递增子序列的最后一个元素的最小值
	tails := []int{}

	for _, height := range heights {
		// 二分查找插入位置
		idx := sort.SearchInts(tails, height)

		if idx == len(tails) {
			// 当前高度大于所有tail中的值，追加到末尾
			tails = append(tails, height)
		} else {
			// 更新对应位置的值为更小的高度
			tails[idx] = height
		}
	}

	return len(tails)
}

func testOne(in string, ans int) {
	{
		envelopes := helper.ParseIntMatrix(in)
		helper.Assert(maxEnvelopes(envelopes) == ans)
	}
	{
		envelopes := helper.ParseIntMatrix(in)
		helper.Assert(maxEnvelopes2(envelopes) == ans)
	}
}

func main() {
	testOne("[[5,4],[6,4],[6,7],[2,3]]", 3)
	testOne("[[1,1],[1,1],[1,1]]", 1)
	testOne("[[4,5],[4,6],[6,7],[2,3],[1,1]]", 4)
}
