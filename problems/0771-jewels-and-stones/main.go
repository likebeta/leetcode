package main

import (
	"leetcode/helper"
)

// 宝石与石头
func numJewelsInStones(J string, S string) (sum int) {
	jewelsSet := map[byte]bool{}
	for i := range J {
		jewelsSet[J[i]] = true
	}
	for i := range S {
		if jewelsSet[S[i]] {
			sum++
		}
	}
	return
}

func main() {
	helper.Log(numJewelsInStones("aA", "aAAbbbb") == 3)
	helper.Log(numJewelsInStones("z", "ZZ") == 0)
}
