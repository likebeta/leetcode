package main

import (
	"leetcode/helper"
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

func main() {
	helper.Log(intersect([]int{1, 2, 2, 1}, []int{2, 2}), []int{2, 2})
	helper.Log(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9})
}
