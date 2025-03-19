package main

import "leetcode/helper"

// 最长公共子序列
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列
// 两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列
// 若这两个字符串没有公共子序列，则返回 0
// 使用动态规划求解，dp[i][j] 表示 text1 的前 i 个字符与 text2 的前 j 个字符的最长公共子序列长度
func longestCommonSubsequence(text1 string, text2 string) int {
	lS, lD := len(text1), len(text2)
	// 处理边界情况：如果其中一个字符串为空，则最长公共子序列长度为0
	if lS == 0 || lD == 0 {
		return 0 // 任一字符串为空，公共子序列长度为0
	}

	// 创建动态规划数组，dp[i][j] 表示 text1[0:i] 与 text2[0:j] 的最长公共子序列长度
	dp := make([][]int, lS+1)
	for i := 0; i < len(dp); i++ {
		// 初始化第一列：与空字符串的最长公共子序列长度为0，默认值已经是0，不需要额外设置
		dp[i] = make([]int, lD+1)
	}

	// 填充dp数组
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// 如果当前字符相同，则最长公共子序列长度加1
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 当前字符不同时，取前面状态的最大值
				// 1. 不使用text1的当前字符：dp[i-1][j]
				// 2. 不使用text2的当前字符：dp[i][j-1]
				// 取这两种情况的最大值
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[lS][lD]
}

// 最长公共子序列的空间优化版本
// 使用一维数组优化空间复杂度，从O(mn)降低到O(min(m,n))
func longestCommonSubsequence2(text1 string, text2 string) int {
	lS, lD := len(text1), len(text2)
	if lS == 0 || lD == 0 {
		return 0
	}

	// 确保text1是较短的字符串，这样可以使用更少的空间
	if lS > lD {
		lS, lD = lD, lS
		text1, text2 = text2, text1
	}

	// 只使用一维数组存储状态
	dp := make([]int, lS+1)

	// 遍历text2的每个字符
	for j := 1; j <= lD; j++ {
		// 保存左上角的值（dp[i-1][j-1]）
		prevDiag := 0
		for i := 1; i <= lS; i++ {
			// 保存当前位置的值，它将成为下一个位置的左上角值
			temp := dp[i]

			// 如果字符相同，则基于左上角值加1
			if text1[i-1] == text2[j-1] {
				dp[i] = prevDiag + 1
			} else {
				// 不相同时，取左边或上边的最大值
				if dp[i-1] > dp[i] {
					dp[i] = dp[i-1]
				}
			}

			// 更新左上角值，为下一个位置做准备
			prevDiag = temp
		}
	}

	return dp[lS]
}

func testOne(text1 string, text2 string, expected int) {
	{
		result := longestCommonSubsequence(text1, text2)
		helper.Assert(result == expected)
	}
	{
		result2 := longestCommonSubsequence2(text1, text2)
		helper.Assert(result2 == expected)
	}
}

func main() {
	// 基本测试用例
	testOne("abcde", "ace", 3)
	testOne("abc", "abc", 3)
	testOne("abc", "def", 0)

	// 边界情况
	testOne("", "", 0)
	testOne("", "abc", 0)
	testOne("abc", "", 0)

	// 相同的字符串
	testOne("abcdef", "abcdef", 6)

	// 部分匹配的情况
	testOne("abcdef", "acf", 3)
	testOne("bsbininm", "jmjkbkjkv", 1)

	// 长度差距较大的情况
	testOne("pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopic", 16)

	// 包含重复字符的情况
	testOne("aaabbb", "ababab", 4)

	// LeetCode官方测试用例
	testOne("ezupkr", "ubmrapg", 2)
	testOne("oxcpqrsvwf", "shmtulqrypy", 2)
}
