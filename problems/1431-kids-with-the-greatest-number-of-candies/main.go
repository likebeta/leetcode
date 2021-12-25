package main

import "leetcode/helper"

// 拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	length := len(candies)
	max := candies[0]
	for i := 1; i < length; i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	ans := make([]bool, length, length)
	for i := range candies {
		if candies[i] + extraCandies >= max {
			ans[i] = true
		}
	}
	return ans
}

func main() {
	helper.Log(kidsWithCandies([]int{2,3,5,1,3}, 3)) // [true,true,true,false,true]
	helper.Log(kidsWithCandies([]int{4,2,1,1,2}, 1)) // [true,false,false,false,false]
	helper.Log(kidsWithCandies([]int{12,1,12}, 10))  // [true,false,true]
}
