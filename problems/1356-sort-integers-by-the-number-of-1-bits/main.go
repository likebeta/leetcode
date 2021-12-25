package main

import (
	"leetcode/helper"
    "sort"
)

// 根据数字二进制下 1 的数目排序
func sortByBits(a []int) []int {
	const Max = 10000
	var bit = [Max + 1]int{}
	for i := 1; i <= Max; i++ {
		bit[i] = bit[i>>1] + i&1
	}
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}

func testOne(in string, ans string) {
	arr := helper.ParseArray(in)
	result := sortByBits(arr)
	helper.Assert(helper.DumpArray(result) == ans)
}

func main() {
	testOne("[0,1,2,3,4,5,6,7,8]", "[0,1,2,4,8,3,5,6,7]")
	testOne("[1024,512,256,128,64,32,16,8,4,2,1]", "[1,2,4,8,16,32,64,128,256,512,1024]")
	testOne("[10000,10000]", "[10000,10000]")
	testOne("[2,3,5,7,11,13,17,19]", "[2,3,5,17,7,11,13,19]")
	testOne("[10,100,1000,10000]", "[10,100,10000,1000]")
}
