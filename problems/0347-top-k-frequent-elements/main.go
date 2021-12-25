package main

import (
	"leetcode/helper"
	"math/rand"
	"time"
)

// 前 K 个高频元素
func topKFrequent(nums []int, k int) []int {
	numCounter := map[int]int{}
	for _, num := range nums {
		numCounter[num]++
	}
	var values [][]int
	for key, value := range numCounter {
		values = append(values, []int{key, value})
	}
	ans := make([]int, k)
	qsort(values, 0, len(values)-1, ans, 0, k)
	return ans
}

func qsort(values [][]int, start, end int, ans []int, ansIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		qsort(values, start, index-1, ans, ansIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ans[ansIndex] = values[i][0]
			ansIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ans, ansIndex, k-(index-start+1))
		}
	}
}

func main() {
	helper.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1, 2})
	helper.Log(topKFrequent([]int{1}, 1), []int{1})
}
