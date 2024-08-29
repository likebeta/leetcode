package main

import (
	"leetcode/helper"
	"sort"
)

// 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx := make([]int, n)
	for i := 1; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] > nums2[idx[j]]
	})

	ans := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[left] > nums2[idx[i]] {
			ans[idx[i]] = nums1[left]
			left++
		} else {
			ans[idx[i]] = nums1[right]
			right--
		}
	}

	return ans
}

func advantageCount2(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx := make([]int, n)
	for i := 1; i < n; i++ {
		idx[i] = i
	}
	sort.Ints(nums1)
	sort.Slice(idx, func(i, j int) bool { return nums2[idx[i]] < nums2[idx[j]] })

	ans := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[i] > nums2[idx[left]] {
			ans[idx[left]] = nums1[i]
			left++
		} else {
			ans[idx[right]] = nums1[i]
			right--
		}
	}
	return ans
}

func testOne(s1, s2 string, out string, f func([]int, []int) []int) {
	arr1 := helper.ParseIntArray(s1)
	arr2 := helper.ParseIntArray(s2)
	outArr := helper.ParseIntArray(out)
	ans := f(arr1, arr2)

	// 答案不唯一
	cnt := 0
	for i := range arr2 {
		if ans[i] > arr2[i] {
			cnt++
		}
	}

	for i := range arr2 {
		if outArr[i] > arr2[i] {
			cnt--
		}
	}

	helper.Assert(cnt == 0)
}

func main() {
	testOne("[2,7,11,15]", "[1,10,4,11]", "[2,11,7,15]", advantageCount)
	testOne("[12,24,8,32]", "[13,25,32,11]", "[24,32,8,12]", advantageCount)
	testOne("[1,5,5,5,5,6,6]", "[0,1,2,3,4,5,6]", "[1,5,5,5,5,6,6]", advantageCount)

	testOne("[2,7,11,15]", "[1,10,4,11]", "[2,11,7,15]", advantageCount2)
	testOne("[12,24,8,32]", "[13,25,32,11]", "[24,32,8,12]", advantageCount2)
	testOne("[1,5,5,5,5,6,6]", "[0,1,2,3,4,5,6]", "[1,5,5,5,5,6,6]", advantageCount2)
}
