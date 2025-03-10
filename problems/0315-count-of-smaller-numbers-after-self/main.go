package main

import (
	"leetcode/helper"
	"sort"
)

// 计算右侧小于当前元素的个数
func countSmaller(nums []int) []int {
	var resultList []int
	c := make([]int, len(nums)+5)
	a := discretization(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(a, nums[i])
		resultList = append(resultList, query(c, id-1))
		update(c, id)
	}
	for i := 0; i < len(resultList)/2; i++ {
		resultList[i], resultList[len(resultList)-1-i] = resultList[len(resultList)-1-i], resultList[i]
	}
	return resultList
}

func lowBit(x int) int {
	return x & (-x)
}

func update(c []int, pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowBit(pos)
	}
}

func query(c []int, pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func discretization(nums []int) []int {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a := make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)
	return a
}

func getId(a []int, x int) int {
	return sort.SearchInts(a, x) + 1
}

type pair struct {
	i int // 下标
	v int // 值
}

func countSmaller2(nums []int) []int {
	n := len(nums)
	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		arr[i] = pair{i, nums[i]}
	}
	tmp := make([]pair, n)
	count := make([]int, n)

	mergeSort(tmp, arr, count, 0, n-1)

	return count
}

func mergeSort(tmp, arr []pair, count []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort(tmp, arr, count, left, mid)
		mergeSort(tmp, arr, count, mid+1, right)
		merge(tmp, arr, count, left, mid, right)
	}
}

func merge(tmp, arr []pair, count []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		tmp[i] = arr[i]
	}

	i, j := left, mid+1
	for idx := left; idx <= right; idx++ {
		if i > mid {
			// 左侧已合并完
			arr[idx], j = tmp[j], j+1
		} else if j > right {
			// 右侧已合并完
			count[tmp[i].i] += j - (mid + 1)
			arr[idx], i = tmp[i], i+1
		} else if tmp[i].v > tmp[j].v {
			// 左侧>右侧
			arr[idx], j = tmp[j], j+1
		} else {
			// 左侧<=右侧
			count[tmp[i].i] += j - (mid + 1)
			arr[idx], i = tmp[i], i+1
		}
	}
}

func testOne(in string, ans string) {
	{
		nums := helper.ParseIntArray(in)
		count := countSmaller(nums)
		helper.Assert(helper.DumpIntArray(count) == ans)
	}

	{
		nums := helper.ParseIntArray(in)
		count := countSmaller2(nums)
		helper.Assert(helper.DumpIntArray(count) == ans)
	}
}

func main() {
	testOne("[5,2,6,1]", "[2,1,1,0]")
}
