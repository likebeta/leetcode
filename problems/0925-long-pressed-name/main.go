package main

import (
	"leetcode/helper"
)

// 长按键入
func isLongPressedName(name string, typed string) bool {
	l, r := len(name), len(typed)
	var i, j int
	for j < r {
		if i < l && name[i] == typed[j] {
			i++
			j++
			continue
		}
		if i == 0 || name[i-1] != typed[j] {
			return false
		}
		j++
	}

	return i == l
}

func main() {
	helper.Assert(isLongPressedName("alex", "aaleex"))
	helper.Assert(!isLongPressedName("saeed", "ssaaedd"))
	helper.Assert(isLongPressedName("leelee", "lleeelee"))
	helper.Assert(isLongPressedName("laiden", "laiden"))
	helper.Assert(!isLongPressedName("hi", "hh"))
	helper.Assert(!isLongPressedName("hi", "h"))
}
