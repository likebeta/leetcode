package main

import (
    "leetcode/helper"
)

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
    var max int
    length := len(s)
    if length > 0 {
        var stat [128]int
        var start int
        for i, v := range s {
            if stat[v] > start {
                if max < i-start {
                    max = i - start
                }
                start = stat[v]
            }
            stat[v] = i + 1
        }
        if max < length-start {
            max = length - start
        }
    }

    return max
}

func main() {
    helper.Assert(3 == lengthOfLongestSubstring("abcabcbb"))
    helper.Assert(1 == lengthOfLongestSubstring("bbbbb"))
    helper.Assert(3 == lengthOfLongestSubstring("pwwkew"))
}
