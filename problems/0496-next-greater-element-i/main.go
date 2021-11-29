package main

import (
	"leetcode/helper"
)

// 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, len(nums2)+1)
	k := 0
	cache := make(map[int]int)

	for i := range nums2 {
		for top := stack[k]; k != 0 && nums2[i] > top; top = stack[k] {
			k--
			cache[top] = nums2[i]
		}
		k++
		stack[k] = nums2[i]
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		if vv, ok := cache[v]; ok {
			ans[i] = vv
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	stack := make([]int, len(nums2)+1)
	k := 0
	stack[k] = -1
	cache := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for top := stack[k]; k != 0 && nums2[i] >= top; top = stack[k] {
			k--
		}
		cache[nums2[i]] = stack[k]
		k++
		stack[k] = nums2[i]
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		ans[i] = cache[v]
	}
	return ans
}

func testOne(in1, in2 string, ans string) {
	var nums1, nums2, result []int
	// 正向
	nums1 = helper.ParseArray(in1)
	nums2 = helper.ParseArray(in2)
	result = nextGreaterElement(nums1, nums2)
	helper.Log(in1, in2, "=1=>", result, ans)
	// 逆向
	nums1 = helper.ParseArray(in1)
	nums2 = helper.ParseArray(in2)
	result = nextGreaterElement2(nums1, nums2)
	helper.Log(in1, in2, "=2=>", result, ans)
}

func main() {
	testOne("[4,1,2]", "[1,3,4,2]", "[-1,3,-1]")
	testOne("[2,4]", "[1,2,3,4]", "[3,-1]")
}
