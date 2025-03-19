package main

import "leetcode/helper"

// 编辑距离
// 给定两个单词 word1 和 word2，计算将 word1 转换成 word2 所使用的最少操作数
// 你可以对一个单词进行如下三种操作：
// 1. 插入一个字符
// 2. 删除一个字符
// 3. 替换一个字符
// 使用动态规划求解，dp[i][j] 表示 word1 的前 i 个字符转换到 word2 的前 j 个字符需要的最小操作数
func minDistance(word1 string, word2 string) int {
	lS, lD := len(word1), len(word2)
	// 处理边界情况：如果其中一个字符串为空，则最小操作数就是另一个字符串的长度
	if lS == 0 {
		return lD // word1为空，需要插入word2的所有字符
	} else if lD == 0 {
		return lS // word2为空，需要删除word1的所有字符
	}

	// 创建动态规划数组，dp[i][j] 表示 word1[0:i] 转换到 word2[0:j] 的最小操作数
	dp := make([][]int, lS+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lD+1)
		dp[i][0] = i // 初始化第一列：word2为空，需要删除word1的字符
	}

	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j // 初始化第一行：word1为空，需要插入word2的字符
	}

	// 填充dp数组，考虑所有可能的编辑操作
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// 如果当前字符相同，则不需要任何操作，直接使用前一个状态
			dp[i][j] = dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				// 当前字符不同时，考虑三种操作：
				// 1. 替换：dp[i-1][j-1] + 1
				// 2. 插入：dp[i][j-1] + 1
				// 3. 删除：dp[i-1][j] + 1
				// 取这三种操作中的最小值
				if dp[i][j] > dp[i][j-1] {
					dp[i][j] = dp[i][j-1] // 插入操作
				}
				if dp[i][j] > dp[i-1][j] {
					dp[i][j] = dp[i-1][j] // 删除操作
				}
				dp[i][j] += 1 // 对选择的操作加1
			}
		}
	}

	return dp[lS][lD]
}

// minDistance2 使用一维数组优化空间复杂度
func minDistance2(word1 string, word2 string) int {
	lS, lD := len(word1), len(word2)
	// 处理边界情况
	if lS == 0 {
		return lD
	} else if lD == 0 {
		return lS
	}

	// 创建一维dp数组，长度为word2的长度+1
	dp := make([]int, lD+1)
	// 初始化dp数组，表示word1为空时的情况
	for j := 0; j <= lD; j++ {
		dp[j] = j
	}

	// 遍历word1的每个字符
	for i := 1; i <= lS; i++ {
		// 保存左上角的值（即dp[i-1][j-1]）
		prev := dp[0]
		// 更新dp[0]，表示当前word1的前i个字符转换到空字符串需要的操作数
		dp[0] = i

		// 遍历word2的每个字符
		for j := 1; j <= lD; j++ {
			// 保存当前dp[j]的值，用于下一次迭代的左上角值
			temp := dp[j]

			// 如果当前字符相同，则使用左上角的值
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				// 当前字符不同时，取三种操作中的最小值
				// dp[j-1]: 插入操作
				// dp[j]: 删除操作
				// prev: 替换操作
				dp[j] = min(min(dp[j-1], dp[j]), prev) + 1
			}

			// 更新左上角的值
			prev = temp
		}
	}

	return dp[lD]
}

func testOne(word1 string, word2 string, expected int) {
	{
		result := minDistance(word1, word2)
		helper.Assert(result == expected)
	}

	{
		result := minDistance2(word1, word2)
		helper.Assert(result == expected)
	}
}

func main() {
	// 基本测试用例
	testOne("horse", "ros", 3)
	testOne("intention", "execution", 5)

	// 边界情况
	testOne("", "", 0)
	testOne("", "abc", 3)
	testOne("abc", "", 3)

	// 相同的字符串
	testOne("abcdef", "abcdef", 0)

	// 只有一个字符不同的情况
	testOne("abcdef", "abcxef", 1)

	// 字符串完全不同的情况
	testOne("abc", "xyz", 3)

	// 部分匹配的情况
	testOne("kitten", "sitting", 3)
	testOne("saturday", "sunday", 3)

	// 长度差距较大的情况
	testOne("pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopic", 29)

	// 包含重复字符的情况
	testOne("aaabbb", "ababab", 2)
}
