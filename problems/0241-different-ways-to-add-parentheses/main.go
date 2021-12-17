package main

import (
	"leetcode/helper"
	"sort"
	"strconv"
)

// 为运算表达式设计优先级
func diffWaysToCompute(expression string) []int {
	var ans []int

	for i, v := range expression {
		if v == '+' || v == '-' || v == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for l := range left {
				for r := range right {
					if v == '+' {
						ans = append(ans, left[l]+right[r])
					} else if v == '-' {
						ans = append(ans, left[l]-right[r])
					} else if v == '*' {
						ans = append(ans, left[l]*right[r])
					}
				}
			}
		}
	}

	if len(ans) == 0 {
		v, _ := strconv.Atoi(expression)
		ans = append(ans, v)
	}

	return ans
}

func testOne(in string, out string) {
	ans := helper.ParseArray(out)
	result := diffWaysToCompute(in)
	sort.Ints(result)
	sort.Ints(ans)
	helper.Log(in, "=>", result)
	helper.Assert(helper.DumpArray(result) == helper.DumpArray(ans))
}

func main() {
	testOne("2-1-1", "[0, 2]")
	testOne("2*3-4*5", "[-34, -14, -10, -10, 10]")
}
