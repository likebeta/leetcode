package main

import (
	"leetcode/helper"
)

// 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	var maxArea int
	length := len(heights)
	if length > 0 {
		var top, tmpArea int
		stack := []int{-1}
		for i := 0; i <= length; i++ {
			for stack[top] != -1 && (i == length || heights[stack[top]] > heights[i]) {
				tmpArea = (i - stack[top-1] - 1) * heights[stack[top]]
				if tmpArea > maxArea {
					maxArea = tmpArea
				}
				top--
			}
			top++
			if len(stack) == top {
				stack = append(stack, i)
			} else {
				stack[top] = i
			}
		}
	}
	return maxArea
}

func main() {
	helper.Assert(largestRectangleArea([]int{2}) == 2)
	helper.Assert(largestRectangleArea([]int{2, 1, 2}) == 3)
	helper.Assert(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}) == 10)
}
