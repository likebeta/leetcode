package main

import (
	"leetcode/helper"
)

// 独一无二的出现次数
func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, v := range arr {
		counter[v]++
	}
	var empty struct{}
	values := make(map[int]struct{})
	for _, v := range counter {
		values[v] = empty
	}
	return len(counter) == len(values)
}

func testOne(in string, ans bool) {
	arr := helper.ParseArray(in)
	helper.Assert(uniqueOccurrences(arr) == ans)
}

func main() {
	testOne("[1,2,2,1,1,3]", true)
	testOne("[1,2]", false)
	testOne("[-3,0,1,-3,1,1,1,-3,10,0]", true)
}
