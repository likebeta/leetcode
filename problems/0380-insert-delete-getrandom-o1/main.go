package main

import (
	"leetcode/helper"
	"math/rand/v2"
)

// RandomizedSet O(1) 时间插入、删除和获取随机元素
type RandomizedSet struct {
	valToKey map[int]int
	values   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToKey: make(map[int]int),
		values:   make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToKey[val]; ok {
		return false
	}

	sz := len(this.values)
	this.values = append(this.values, val)
	this.valToKey[val] = sz
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.valToKey[val]; ok {
		last := len(this.values) - 1
		lastV := this.values[last]
		this.valToKey[lastV] = idx
		this.values[idx] = lastV

		delete(this.valToKey, val)
		this.values = this.values[:last]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	sz := len(this.values)
	n := rand.IntN(sz)
	return this.values[n]
}

func main() {
	randomizedSet := Constructor()                  // Constructor
	helper.Assert(randomizedSet.Insert(1) == true)  // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	helper.Assert(randomizedSet.Remove(2) == false) // 返回 false ，表示集合中不存在 2 。
	helper.Assert(randomizedSet.Insert(2) == true)  // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	n := randomizedSet.GetRandom()                  // getRandom 应随机返回 1 或 2 。
	helper.Assert(n == 1 || n == 2)                 // 1 or 2
	helper.Assert(randomizedSet.Remove(1) == true)  // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	helper.Assert(randomizedSet.Insert(2) == false) // 2 已在集合中，所以返回 false 。
	n = randomizedSet.GetRandom()                   // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
	helper.Assert(n == 2)                           // 2
}
