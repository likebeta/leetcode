package main

import (
	"leetcode/helper"
	"math/rand"
)

// 排序数组
func sortArray(nums []int) []int {
	doQuicksort(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
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
// 最终停止的条件是left==right，并且left处是等待填充的坑，解释如下：
//  1. 外层for循环开始时，left处是上次或者初始化留下的坑
//  2. 如果right左移导致left==right，此时left不再右移，left处依然是坑
//  3. 如果left右移导致left==right，right所在值已填入右移前left处，当前是坑，而left==right，所以left处也是坑
//  4. 如果right左移和left右移后left<right，left处先填，right后填，回到1
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
// 最终停止的条件是left==right，并且left处<=arr[pivot]，解释如下：
//  1. 外层for循环开始时，left处是上次交换后或者初始化 <=arr[pivot]
//  2. 如果right左移导致left==right，此时left不再右移，left处保持不变依然 <= arr[pivot]
//  3. 如果left右移导致left==right，由于right结束条件是<arr[pivot]，而left==right，所以left处也<arr[poivt]
//  4. 如果right左移和left右移后left<right，交换后，left处是交换前right处值(结束条件是<arr[pivot]), 回到1
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
// 最终停止的条件是i==left, 此时 (left,pivot] <= arr[left] < (pivot,right]
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

func sortArray2(nums []int) []int {
	tmp := make([]int, len(nums))
	mergeSort(tmp, nums, 0, len(nums)-1)
	return nums
}

// 归并排序
func mergeSort(tmp, nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort(tmp, nums, left, mid)
		mergeSort(tmp, nums, mid+1, right)
		merge(tmp, nums, left, mid, right)
	}
}

func merge(tmp, nums []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		tmp[i] = nums[i]
	}

	i, j := left, mid+1
	for idx := left; idx <= right; idx++ {
		if i > mid {
			// 左侧已合并完
			nums[idx], j = tmp[j], j+1
		} else if j > right {
			// 右侧已合并完
			nums[idx], i = tmp[i], i+1
		} else if tmp[i] > tmp[j] {
			// 左侧>右侧
			nums[idx], j = tmp[j], j+1
		} else {
			// 左侧<=右侧
			nums[idx], i = tmp[i], i+1
		}
	}
}

func testOne(in string, ans string) {
	{
		arr := helper.ParseArray(in)
		sortArray(arr)
		result := helper.DumpArray(arr)
		helper.Assert(result == ans)
	}
	{
		arr := helper.ParseArray(in)
		sortArray2(arr)
		result := helper.DumpArray(arr)
		helper.Assert(result == ans)
	}
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
