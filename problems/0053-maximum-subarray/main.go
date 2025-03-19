package main

import (
	"leetcode/helper"
	"math"
)

// 最大子数组和 - 基础DP解法
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// 最大子数组和 - O(n)时间，O(1)空间
func maxSubArray2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 使用两个变量替代整个dp数组
	// currMax表示以当前元素结尾的最大子数组和
	// globalMax表示全局最大子数组和
	currMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < n; i++ {
		// 如果currMax大于0，则继续累加；否则从当前元素重新开始
		currMax = max(nums[i], currMax+nums[i])
		globalMax = max(globalMax, currMax)
	}

	return globalMax
}

// 暴力解法 - O(n²)
func maxSubArray3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxSum := math.MinInt32
	// 枚举所有可能的子数组
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}

// 前缀和解法 - O(n)
func maxSubArray4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 初始化
	maxSum := nums[0]
	prefixSum := 0
	minPrefixSum := 0

	for _, num := range nums {
		// 更新前缀和
		prefixSum += num

		// 当前最大子数组和 = 当前前缀和 - 之前的最小前缀和
		maxSum = max(maxSum, prefixSum-minPrefixSum)

		// 更新最小前缀和（用于后续计算）
		minPrefixSum = min(minPrefixSum, prefixSum)
	}

	return maxSum
}

func testOne(in string, ans int) {
	{
		nums := helper.ParseIntArray(in)
		helper.Assert(maxSubArray(nums) == ans)
	}
	{
		nums := helper.ParseIntArray(in)
		helper.Assert(maxSubArray2(nums) == ans)
	}
	{
		nums := helper.ParseIntArray(in)
		helper.Assert(maxSubArray3(nums) == ans)
	}
	{
		nums := helper.ParseIntArray(in)
		helper.Assert(maxSubArray4(nums) == ans)
	}
}

func main() {
	testOne("[-2,1,-3,4,-1,2,1,-5,4]", 6)
	testOne("[1]", 1)
	testOne("[5,4,-1,7,8]", 23)
}
