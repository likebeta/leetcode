package main

import (
	"leetcode/helper"
)

// 最小栈
type MinStack struct {
	dataArr []int
	minArr  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.dataArr = append(this.dataArr, x)
	length := len(this.minArr)
	if length == 0 || this.minArr[length-1] >= x {
		this.minArr = append(this.minArr, x)
	}
}

func (this *MinStack) Pop() {
	ld := len(this.dataArr)
	if ld > 0 {
		lm := len(this.minArr)
		if this.minArr[lm-1] == this.dataArr[ld-1] {
			this.minArr = this.minArr[:lm-1]
		}
		this.dataArr = this.dataArr[:ld-1]
	}
}

func (this *MinStack) Top() int {
	return this.dataArr[len(this.dataArr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	helper.Assert(minStack.GetMin() == -3) // --> 返回 -3.
	minStack.Pop()
	helper.Assert(minStack.Top() == 0)     // --> 返回 0.
	helper.Assert(minStack.GetMin() == -2) // --> 返回 -2.
}
