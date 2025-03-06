package main

import (
	"leetcode/helper"
	"math/rand"
)

// 黑名单中的随机数
// b2w:   [0, bound)中黑名单数字映射到[bound, n)中有效数字
// bound: 有效数字范围 = 总数量 - 黑名单数量
type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m
	black := map[int]bool{}
	for _, b := range blacklist {
		if b >= bound {
			black[b] = true
		}
	}

	b2w := make(map[int]int, m-len(black))
	w := bound
	for _, b := range blacklist {
		if b < bound {
			for black[w] {
				w++
			}
			b2w[b] = w
			w++
		}
	}
	return Solution{b2w, bound}
}

func (s *Solution) Pick() int {
	x := rand.Intn(s.bound)
	if s.b2w[x] > 0 {
		return s.b2w[x]
	}
	return x
}

func main() {
	solution := Constructor(7, []int{2, 3, 5})
	helper.LogF("pick=%d", solution.Pick())
	helper.LogF("pick=%d", solution.Pick())
	helper.LogF("pick=%d", solution.Pick())
	helper.LogF("pick=%d", solution.Pick())
}
