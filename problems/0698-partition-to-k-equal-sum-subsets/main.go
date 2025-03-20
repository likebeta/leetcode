package main

import (
	"leetcode/helper"
	"slices"
)

// 划分为k个相等的子集
// 判断是否可以将数组划分为k个和相等的子集
// 实现思路：以桶的视角，每次尝试将数字放入当前桶中，直到所有桶都填满
func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}

	target := sum / k

	// 优化：对数组进行排序和反转，使大数在前面
	slices.Sort(nums)
	slices.Reverse(nums)

	// 如果最大的数大于目标和，直接返回false
	if nums[0] > target {
		return false
	}

	used := make([]bool, n)

	return backtrackPartition(nums, used, k, target, n, 0, 0, 0)
}

// 回溯函数，尝试将数字放入当前子集中
//   - nums: 输入的整数数组
//   - used: 标记数组，记录每个数字是否已被使用
//   - k: 要划分的子集数量
//   - target: 每个子集的目标和
//   - n: 数组长度
//   - index: 当前考虑的数字索引
//   - count: 已经找到的满足条件的子集数量
//   - curSum: 当前子集的和
func backtrackPartition(nums []int, used []bool, k, target, n, index, count, curSum int) bool {
	// 已经找到k-1个子集，最后一个子集一定满足条件
	if count == k-1 {
		return true
	}

	// 当前子集已满足目标和，开始下一个子集
	if curSum == target {
		return backtrackPartition(nums, used, k, target, n, 0, count+1, 0)
	}

	// 尝试将每个未使用的数加入当前子集
	for i := index; i < n; i++ {
		if used[i] || curSum+nums[i] > target {
			continue
		}

		used[i] = true
		if backtrackPartition(nums, used, k, target, n, i+1, count, curSum+nums[i]) {
			return true
		}
		used[i] = false

		// 优化：如果当前和为0且失败，则后续都会失败
		if curSum == 0 {
			break
		}
	}

	return false
}

// 判断是否可以将数组划分为k个和相等的子集
// 实现思路：以球的视角，每次尝试将当前数字放入某个桶中，直到所有数字都被放置
func canPartitionKSubsets2(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}

	target := sum / k

	// 优化：对数组进行排序和反转，使大数在前面
	slices.Sort(nums)
	slices.Reverse(nums)

	// 如果最大的数大于目标和，直接返回false
	if nums[0] > target {
		return false
	}

	// 记录每个桶的当前和
	buckets := make([]int, k)

	return backtrackPartition2(nums, buckets, target, 0)
}

// 回溯函数，尝试将当前数字放入不同的桶中
//   - nums: 输入的整数数组
//   - buckets: 记录每个桶当前的和
//   - target: 每个桶的目标和
//   - index: 当前要放置的数字的索引
func backtrackPartition2(nums []int, buckets []int, target, index int) bool {
	// 所有数字都已经放入桶中
	if index == len(nums) {
		return true
	}

	// 尝试将当前数字放入每个桶
	for i := 0; i < len(buckets); i++ {
		// 剪枝：如果当前桶和前一个桶的和相同，跳过，因为放入前一个桶和当前桶的结果是一样的
		if i > 0 && buckets[i] == buckets[i-1] {
			continue
		}

		// 如果当前桶加上当前数字不超过目标和
		if buckets[i]+nums[index] <= target {
			// 将当前数字放入当前桶
			buckets[i] += nums[index]
			// 继续放置下一个数字
			if backtrackPartition2(nums, buckets, target, index+1) {
				return true
			}
			// 回溯：从当前桶中移除数字
			buckets[i] -= nums[index]
		}
	}

	return false
}

func testOne(in string, k int, ans bool) {
	{
		nums := helper.ParseIntArray(in)
		helper.Assert(canPartitionKSubsets(nums, k) == ans)
	}

	{
		nums := helper.ParseIntArray(in)
		helper.Assert(canPartitionKSubsets2(nums, k) == ans)
	}
}

func main() {
	// 基本测试用例
	testOne("[4,3,2,3,5,2,1]", 4, true) // 可以分成4个子集，每个子集和为5
	testOne("[1,2,3,4]", 3, false)      // 总和为10，不能被3整除

	// 边界情况
	testOne("[1]", 1, true)              // k=1时总是可以分割
	testOne("[1,1]", 2, true)            // k等于数组长度
	testOne("[2,2,2,2,3,4,5]", 4, false) // 总和可以被k整除，但无法分割

	// 复杂测试用例
	testOne("[10,10,10,7,7,7,7,7,7,6,6,6]", 3, true)     // 可以分成3个子集，每个子集和为30
	testOne("[2,2,2,2,2,2,2,2,2,2,2,2,2,3,3]", 8, false) // 总和可以被k整除，但无法分割

	// 特殊情况
	testOne("[1,1,1,1,2,2,2,2]", 2, true)         // 多个相同的数字
	testOne("[5,5,5,5,4,4,4,4,3,3,3,3]", 3, true) // 大数在前面的情况
}
