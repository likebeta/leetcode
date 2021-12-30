package main

import (
	"github.com/likebeta/leetcode/helper"
)

// 统计特殊四元组
func countQuadruplets(nums []int) (ans int) {
	counter := make(map[int]int)
	for k := len(nums) - 2; k >= 2; k-- {
		counter[nums[k+1]]++
		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				sum := nums[k] + nums[i] + nums[j]
				ans += counter[sum]
			}
		}
	}
	return
}

func countQuadruplets2(nums []int) (ans int) {
	counter := make(map[int]int)
	for k := len(nums) - 3; k >= 1; k-- {
		for i := k + 2; i < len(nums); i++ {
			counter[nums[i]-nums[k+1]]++
		}
		for i := 0; i < k; i++ {
			sum := nums[k] + nums[i]
			ans += counter[sum]
		}
	}
	return
}

func testOne(in string, ans int) {
	nums := helper.ParseArray(in)
	helper.Assert(countQuadruplets(nums) == ans)
	helper.Assert(countQuadruplets2(nums) == ans)
}

func main() {
	testOne("[1,2,3,6]", 1)
	testOne("[3,3,6,4,5]", 0)
	testOne("[1,1,1,3,5]", 4)
}
