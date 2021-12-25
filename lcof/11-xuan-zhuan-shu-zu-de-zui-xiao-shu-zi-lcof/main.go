package main

import "leetcode/helper"

// 旋转数组的最小数字
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := i + (j-i)/2
		if numbers[mid] > numbers[j] {
			i = mid+1
		} else if numbers[mid] < numbers[j] {
			j = mid
		} else {
			j--
		}
	}
	return numbers[i]
}

func main() {
	helper.Assert(minArray([]int{3, 4, 5}) == 3)
	helper.Assert(minArray([]int{3, 4, 5, 1}) == 1)
	helper.Assert(minArray([]int{3, 4, 5, 1, 2}) == 1)
	helper.Assert(minArray([]int{2, 2, 2, 0, 1}) == 0)
}
