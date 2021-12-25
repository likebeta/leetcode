package main

import (
	"leetcode/helper"
)

func sortArrayByParityII(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

func testOne(in string) {
	arr := helper.ParseArray(in)
	result := sortArrayByParityII(arr)
	helper.Log(in, helper.DumpArray(result))
}

func main() {
	testOne("[4,2,5,7]")
}
