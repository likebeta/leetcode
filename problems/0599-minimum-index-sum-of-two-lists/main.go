package main

import (
	"leetcode/helper"
)

// 两个列表的最小索引总和
func findRestaurant(list1 []string, list2 []string) []string {
	cache := make(map[string]int, len(list1))
	for i, v := range list1 {
		cache[v] = i
	}
	var ans []string
	indexSum := 1<<31 - 1
	for j, v := range list2 {
		if i, ok := cache[v]; ok {
			if indexSum > i+j {
				indexSum = i + j
				ans = []string{v}
			} else if indexSum == i+j {
				ans = append(ans, v)
			}
		}
	}
	return ans
}

func testOne(in1, in2 string, result string) {
	list1 := helper.ParseStrArray(in1)
	list2 := helper.ParseStrArray(in2)
	resultArr := helper.ParseStrArray(result)
	ans := findRestaurant(list1, list2)
	helper.Assert(helper.Dump(ans) == helper.Dump(resultArr))
}

func main() {
	testOne(`["Shogun", "Tapioca Express", "Burger King", "KFC"]`, `["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]`, `["Shogun"]`)
	testOne(`["Shogun", "Tapioca Express", "Burger King", "KFC"]`, `["KFC", "Shogun", "Burger King"]`, `["Shogun"]`)
}
