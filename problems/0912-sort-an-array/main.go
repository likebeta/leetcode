package main

import (
	"github.com/likebeta/leetcode/helper"
	"math/rand"
)

// 快速排序
func sortArray(nums []int) []int {
	doQuicksort(nums, 0, len(nums)-1)
	return nums
}

func doQuicksort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		doQuicksort(arr, left, pivot-1)
		doQuicksort(arr, pivot+1, right)
	}
}

func shufflePivot(arr []int, left, right int) {
	i := rand.Intn(right-left+1) + left
	arr[i], arr[left] = arr[left], arr[i]
}

// 挖坑法
func partition(arr []int, left, right int) int {
	shufflePivot(arr, left, right)
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

// 交换法
func partition2(arr []int, left, right int) int {
	shufflePivot(arr, left, right)
	pivot := left
	for left < right {
		for left < right && arr[right] >= arr[pivot] {
			right--
		}
		for left < right && arr[left] <= arr[pivot] {
			left++
		}
		arr[right], arr[left] = arr[left], arr[right]
	}
	arr[left], arr[pivot] = arr[pivot], arr[left]
	return left
}

// 顺序遍历法
func partition3(arr []int, left, right int) int {
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

func testOne(in string, ans string) {
	arr := helper.ParseArray(in)
	sortArray(arr)
	result := helper.DumpArray(arr)
	helper.Assert(result == ans)
}

func main() {
	testOne("[]", "[]")
	testOne("[1]", "[1]")
	testOne("[1,3]", "[1,3]")
	testOne("[1,2,3]", "[1,2,3]")
	testOne("[1,5,3]", "[1,3,5]")
	testOne("[5,2,3,1]", "[1,2,3,5]")
	testOne("[1,5,3,6]", "[1,3,5,6]")
	testOne("[1,3,4,3,5,2]", "[1,2,3,3,4,5]")
}
