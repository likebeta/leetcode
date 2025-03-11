package main

import (
	"leetcode/helper"
	"math/rand"
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

func findKthLargest3(nums []int, k int) int {
	kSmallest := len(nums) - k

	left, right := 0, len(nums)-1
	for left <= right {
		pivot := partition(nums, left, right)
		if pivot == kSmallest {
			return nums[pivot]
		} else if pivot > kSmallest {
			right = pivot - 1
		} else { // pivot < kSmallest
			left = pivot + 1
		}
	}
	return -1
}

func shufflePivot(arr []int, left, right int) {
	i := rand.Intn(right-left+1) + left
	arr[i], arr[left] = arr[left], arr[i]
}

func partition(arr []int, left, right int) int {
	shufflePivot(arr, left, right)
	pivot := right
	for i := right; i > left; i-- {
		if arr[i] > arr[left] {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot--
		}
	}
	arr[pivot], arr[left] = arr[left], arr[pivot]
	return pivot
}

func testOne(in string, k int, ans int) {
	{
		arr := helper.ParseIntArray(in)
		helper.Assert(findKthLargest(arr, k) == ans)
	}

	{
		arr := helper.ParseIntArray(in)
		helper.Assert(findKthLargest2(arr, k) == ans)
	}

	{
		arr := helper.ParseIntArray(in)
		helper.Assert(findKthLargest3(arr, k) == ans)
	}
}

func main() {
	testOne("[3, 2, 1, 5, 6, 4]", 5, 2)
	testOne("[3, 2, 3, 1, 2, 4, 5, 5, 6]", 4, 4)
}
