package main

import "leetcode/helper"

// 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}

func main() {
	helper.Log(twoSum([]int{2, 3, 4}, 6), []int{1, 3})
	helper.Log(twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	helper.Log(twoSum([]int{0, 0, 3, 4}, 0), []int{1, 2})

}
