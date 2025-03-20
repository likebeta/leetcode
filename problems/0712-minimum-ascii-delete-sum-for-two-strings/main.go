package main

import (
	"leetcode/helper"
)

// 最小ASCII删除和
// 给定两个字符串s1和s2，返回使两个字符串相等所需删除字符的最小ASCII和。
// 使用动态规划求解，dp[i][j] 表示使 s1 的前 i 个字符和 s2 的前 j 个字符相同所需删除字符的最小ASCII和
func minimumDeleteSum(s1 string, s2 string) int {
	lS, lD := len(s1), len(s2)
	// 处理边界情况：如果其中一个字符串为空，则需要删除另一个字符串的所有字符
	if lS == 0 {
		// 计算s2所有字符的ASCII和
		sum := 0
		for i := 0; i < lD; i++ {
			sum += int(s2[i])
		}
		return sum
	} else if lD == 0 {
		// 计算s1所有字符的ASCII和
		sum := 0
		for i := 0; i < lS; i++ {
			sum += int(s1[i])
		}
		return sum
	}

	// 创建动态规划数组，dp[i][j] 表示使 s1[0:i] 和 s2[0:j] 相同所需删除字符的最小ASCII和
	dp := make([][]int, lS+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lD+1)
	}

	// 初始化第一行和第一列
	// 第一列：s2为空，需要删除s1的所有字符
	for i := 1; i <= lS; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	// 第一行：s1为空，需要删除s2的所有字符
	for j := 1; j <= lD; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	// 填充dp数组
	for i := 1; i <= lS; i++ {
		for j := 1; j <= lD; j++ {
			// 如果当前字符相同，则不需要删除操作，直接使用前一个状态
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 当前字符不同时，可以选择：
				// 1. 删除s1的当前字符：dp[i-1][j] + s1[i-1]的ASCII值
				// 2. 删除s2的当前字符：dp[i][j-1] + s2[j-1]的ASCII值
				// 取这两种操作中的最小值
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[lS][lD]
}

// minimumDeleteSum2 使用一维数组优化空间复杂度
func minimumDeleteSum2(s1 string, s2 string) int {
	lS, lD := len(s1), len(s2)
	// 处理边界情况
	if lS == 0 {
		sum := 0
		for i := 0; i < lD; i++ {
			sum += int(s2[i])
		}
		return sum
	} else if lD == 0 {
		sum := 0
		for i := 0; i < lS; i++ {
			sum += int(s1[i])
		}
		return sum
	}

	// 创建一维dp数组，长度为s2的长度+1
	dp := make([]int, lD+1)
	// 初始化dp数组，表示s1为空时的情况
	for j := 1; j <= lD; j++ {
		dp[j] = dp[j-1] + int(s2[j-1])
	}

	// 遍历s1的每个字符
	for i := 1; i <= lS; i++ {
		// 保存左上角的值（即dp[i-1][j-1]）
		prev := dp[0]
		// 更新dp[0]，表示当前s1的前i个字符和空字符串所需的ASCII和
		dp[0] += int(s1[i-1])

		// 遍历s2的每个字符
		for j := 1; j <= lD; j++ {
			// 保存当前dp[j]的值，用于下一次迭代的左上角值
			temp := dp[j]

			// 如果当前字符相同，则使用左上角的值
			if s1[i-1] == s2[j-1] {
				dp[j] = prev
			} else {
				// 当前字符不同时，取两种删除操作中的最小值
				// dp[j-1] + s2[j-1]: 删除s2当前字符
				// dp[j] + s1[i-1]: 删除s1当前字符
				dp[j] = min(dp[j-1]+int(s2[j-1]), dp[j]+int(s1[i-1]))
			}

			// 更新左上角的值
			prev = temp
		}
	}

	return dp[lD]
}

func testOne(s1 string, s2 string, expected int) {
	{
		result := minimumDeleteSum(s1, s2)
		helper.Assert(result == expected)
	}

	{
		result := minimumDeleteSum2(s1, s2)
		helper.Assert(result == expected)
	}
}

func main() {
	// 基本测试用例
	testOne("sea", "eat", 231)     // 删除's'和't'，ASCII和为115+116=231
	testOne("delete", "leet", 403) // 删除'd'和'e'，ASCII和为100+101+101+101=403

	// 边界情况
	testOne("", "", 0)
	testOne("", "abc", 294) // 删除"abc"，ASCII和为97+98+99=294
	testOne("abc", "", 294) // 删除"abc"，ASCII和为97+98+99=294

	// 相同的字符串
	testOne("abcdef", "abcdef", 0)

	// 完全不同的字符串
	testOne("abc", "def", 597) // 删除所有字符，ASCII和为97+98+99+100+101+102=597

	// 部分匹配的情况
	testOne("abcde", "ace", 198)           // 删除'b'和'd'，ASCII和为98+100=198
	testOne("intention", "execution", 878) // 删除'i','n','t','n'和'x','c','u','t'，ASCII和=105+110+116+110+120+99+117+101=878
}
