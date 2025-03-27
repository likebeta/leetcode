package main

import (
	"leetcode/helper"
)

// 分割等和子集
// 给定一个只包含正整数的非空数组，判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
func canPartition(nums []int) bool {
	n := len(nums)
	// 如果数组长度小于2，无法分割成两个非空子集
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	// 计算数组元素总和和最大值
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 如果总和为奇数，无法分割成两个和相等的子集
	if sum%2 != 0 {
		return false
	}

	// 目标和为总和的一半
	target := sum / 2
	// 如果最大值大于目标和，无法分割成两个和相等的子集
	if max > target {
		return false
	}

	// 二维动态规划解法
	// dp[i][j]表示前i个数字能否通过取舍组成和为j
	// dp[n][target]表示给定的数组能否通过取舍组成和为target
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 初始化：所有行的第0列都为true，表示和为0的子集始终可以通过不选择任何元素得到
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	// 第一个数字可以单独组成和为nums[0]的子集
	dp[0][nums[0]] = true
	// 填充dp数组
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				// 当前数字可以放入背包时，可以选择放入或不放入
				// dp[i-1][j] 表示不放入当前数字，dp[i-1][j-v] 表示放入当前数字
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				// 当前数字大于剩余容量，只能不放入
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	// 返回是否能组成和为target的子集
	return dp[n-1][target]
}

// 一维动态规划优化版本
func canPartition2(nums []int) bool {
	n := len(nums)
	// 如果数组长度小于2，无法分割成两个非空子集
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	// 计算数组元素总和和最大值
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 如果总和为奇数，无法分割成两个和相等的子集
	if sum%2 != 0 {
		return false
	}

	// 目标和为总和的一半
	target := sum / 2
	// 如果最大值大于目标和，无法分割成两个和相等的子集
	if max > target {
		return false
	}

	// 一维动态规划，dp[j]表示是否可以组成和为j的子集
	dp := make([]bool, target+1)
	// 和为0的子集始终可行（不选择任何元素）
	dp[0] = true
	// 遍历每个数字
	for i := 0; i < n; i++ {
		v := nums[i]
		// 注意这里要从大到小遍历，避免一个数字被重复使用
		for j := target; j >= v; j-- {
			// dp[j]表示不放入当前数字，dp[j-v]表示放入当前数字
			dp[j] = dp[j] || dp[j-v]
		}
	}
	// 返回是否能组成和为target的子集
	return dp[target]
}

// 测试函数，验证两个算法是否都返回正确结果
func testOne(arr []int, ans bool) {
	helper.Assert(canPartition(arr) == ans)
	helper.Assert(canPartition2(arr) == ans)
}

func main() {
	// 测试用例1：[1,5,11,5]可以分割成[1,5,5]和[11]，和都为11
	testOne([]int{1, 5, 11, 5}, true)
	// 测试用例2：[1,2,3,5]无法分割成两个和相等的子集
	testOne([]int{1, 2, 3, 5}, false)
}
