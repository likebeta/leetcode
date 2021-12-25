package main

import (
	"leetcode/helper"
	"sort"
)

// 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	N1, N2 := len(nums1), len(nums2)
	if N1 > N2 {
		return intersect(nums2, nums1)
	}
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}
	var ans []int
	for _, num := range nums2 {
		if nums1Map[num] > 0 {
			ans = append(ans, num)
			nums1Map[num]--
		}
	}
	return ans
}

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	N1, N2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	var ans []int
	for index1 < N1 && index2 < N2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			ans = append(ans, nums1[index1])
			index1++
			index2++
		}
	}
	return ans
}

func testOne(arr1, arr2, ans []int) {
	helper.Log(intersect(arr1, arr2), ans)
	helper.Log(intersect2(arr1, arr2), ans)
}

func main() {
	testOne([]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2})
	testOne([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9})
}
