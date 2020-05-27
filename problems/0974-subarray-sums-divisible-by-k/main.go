package main

import (
	"leetcode/helper"
)

// 和可被 K 整除的子数组
func subarraysDivByK(A []int, K int) int {
	cache := map[int]int{0: 1}
	var ans, sum int
	for i := range A {
		sum += A[i]
		mod := sum % K
		// if mod == 0 {
		// 	ans += cache[0]
		// } else if mod > 0 {
		// 	ans += cache[mod] + cache[mod-K]
		// } else {
		// 	ans += cache[mod] + cache[mod+K]
		// }
		ans += cache[mod] + cache[mod-K] + cache[mod+K]
		cache[mod]++
	}
	return ans
}

func subarraysDivByK2(A []int, K int) int {
	cache := map[int]int{0: 1}
	var ans, sum int
	for i := range A {
		sum += A[i]
		mod := (sum%K + K) % K
		ans += cache[mod]
		cache[mod]++
	}
	return ans
}

func main() {
	var arr []int
	arr = []int{4, 5, 0, -2, -3, 1}
	helper.Log(subarraysDivByK(arr, 5), subarraysDivByK2(arr, 5), 7)
	arr = []int{4, 5, 0, 0, -2, -3, 1}
	helper.Log(subarraysDivByK(arr, 5), subarraysDivByK2(arr, 5), 11)
	arr = []int{-1, 2, 9}
	helper.Log(subarraysDivByK(arr, 2), subarraysDivByK2(arr, 2), 2)
	arr = []int{2, -2, 2, -4}
	helper.Log(subarraysDivByK(arr, 6), subarraysDivByK2(arr, 6), 2)
	arr = []int{7, -5, 5, -8, -6, 6, -4, 7, -8, -7}
	helper.Log(subarraysDivByK(arr, 7), subarraysDivByK2(arr, 7), 11)
}
