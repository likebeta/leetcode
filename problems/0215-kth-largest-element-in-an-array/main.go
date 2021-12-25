package main

import (
	"leetcode/helper"
)

// 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
	for i := 1; i < k; i++ {
		heapSize--
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

func findKthLargest2(nums []int, k int) int {
	for i := k / 2; i >= 0; i-- {
		minHeapify(nums, i, k)
	}
	N := len(nums)
	for i := k; i < N; i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			minHeapify(nums, 0, k)
		}
	}
	return nums[0]
}

func minHeapify(a []int, i, heapSize int) {
	l, r, min := i*2+1, i*2+2, i
	if l < heapSize && a[l] < a[min] {
		min = l
	}
	if r < heapSize && a[r] < a[min] {
		min = r
	}
	if min != i {
		a[i], a[min] = a[min], a[i]
		minHeapify(a, min, heapSize)
	}
}

func testOne(arr []int, k int, ans int) {
	helper.Log(findKthLargest(arr, k) == ans)
	helper.Log(findKthLargest2(arr, k) == ans)
}

func main() {
	testOne([]int{3, 2, 1, 5, 6, 4}, 5, 2)
	testOne([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4)
}
