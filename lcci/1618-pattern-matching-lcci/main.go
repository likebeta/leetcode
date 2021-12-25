package main

import (
	"leetcode/helper"
)

// 模式匹配
func patternMatching(pattern string, value string) bool {
	var (
		lP     = len(pattern)
		lV     = len(value)
		counts [2]int
		chars  = [2]byte{'a', 'b'}
	)

	for i := 0; i < lP; i++ {
		if pattern[i] == chars[0] {
			counts[0]++
		} else {
			counts[1]++
		}
	}

	if counts[0] < counts[1] {
		counts[0], counts[1] = counts[1], counts[0]
		chars[0], chars[1] = chars[1], chars[0]
	}

	if lV == 0 {
		return counts[1] == 0
	}
	if lP == 0 {
		return false
	}

	for lLong := 0; counts[0]*lLong <= lV; lLong++ {
		rest := lV - counts[0]*lLong
		if (counts[1] == 0 && rest == 0) || (counts[1] != 0 && rest%counts[1] == 0) {
			var lShort int
			if counts[1] != 0 {
				lShort = rest / counts[1]
			}
			pos, correct := 0, true
			var vLong, vShort string
			for i := 0; i < lP; i++ {
				if pattern[i] == chars[0] {
					sub := value[pos : pos+lLong]
					if len(vLong) == 0 {
						vLong = sub
					} else if vLong != sub {
						correct = false
						break
					}
					pos += lLong
				} else {
					sub := value[pos : pos+lShort]
					if len(vShort) == 0 {
						vShort = sub
					} else if vShort != sub {
						correct = false
						break
					}
					pos += lShort
				}
			}
			if correct && vLong != vShort {
				return true
			}
		}
	}
	return false
}

func main() {
	helper.Assert(patternMatching("a", "") == true)
	helper.Assert(patternMatching("abba", "dogcatcatdog") == true)
	helper.Assert(patternMatching("abba", "dogcatcatfish") == false)
	helper.Assert(patternMatching("aaaa", "dogcatcatdog") == false)
	helper.Assert(patternMatching("abba", "dogdogdogdog") == true)
	helper.Assert(patternMatching("abba", "dogcatcatdog") == true)
}
