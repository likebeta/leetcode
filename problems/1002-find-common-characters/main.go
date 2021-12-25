package main

import (
	"leetcode/helper"
	"math"
)

// 查找常用字符
func commonChars(a []string) (ans []string) {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range a {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}

func main() {
	helper.Log(commonChars([]string{"bella", "label", "roller"}))
	helper.Log(commonChars([]string{"cool", "lock", "cook"}))
}
