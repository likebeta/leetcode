package main

import (
	"leetcode/helper"
)

// 最大间距
func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

type pair struct{ min, max int }

func maximumGap2(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	minVal := min(nums...)
	maxVal := max(nums...)
	d := max(1, (maxVal-minVal)/(n-1))
	bucketSize := (maxVal-minVal)/d + 1

	// 存储 (桶内最小值，桶内最大值) 对，(-1, -1) 表示该桶是空的
	buckets := make([]pair, bucketSize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	for _, v := range nums {
		bid := (v - minVal) / d
		if buckets[bid].min == -1 {
			buckets[bid].min = v
			buckets[bid].max = v
		} else {
			buckets[bid].min = min(buckets[bid].min, v)
			buckets[bid].max = max(buckets[bid].max, v)
		}
	}

	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			ans = max(ans, b.min-buckets[prev].max)
		}
		prev = i
	}
	return
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func testOne(in string, ans int) {
	arr := helper.ParseArray(in)
	helper.Assert(maximumGap(arr) == ans)
	helper.Assert(maximumGap2(arr) == ans)
}

func main() {
	testOne("[3,6,9,1]", 3)
	testOne("[10]", 0)
}
