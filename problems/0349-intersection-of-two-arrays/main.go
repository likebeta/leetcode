package main

import (
	"leetcode/helper"
)

// 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	N1, N2 := len(nums1), len(nums2)
	if N1 > N2 {
		return intersection(nums2, nums1)
	}
	nums1Set := make(map[int]bool)
	for _, v := range nums1 {
		nums1Set[v] = true
	}
	intersectSet := make(map[int]bool)
	for _, v := range nums2 {
		if nums1Set[v] {
			intersectSet[v] = true
		}
	}
	var ans []int
	for k := range intersectSet {
		ans = append(ans, k)
	}
	return ans
}

func testOne(in1, in2 string, ans string) {
	nums1 := helper.ParseArray(in1)
	nums2 := helper.ParseArray(in2)
	result := intersection(nums1, nums2)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("[1,2,2,1]", "[2,2]", "[2]")
	testOne("[4,9,5]", "[9,4,9,8,4]", "[9,4]")
	testOne("[2,6,2,9,1]", "[7,1]", "[1]")
}
