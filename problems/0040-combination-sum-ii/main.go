package main

import (
	"leetcode/helper"
	"slices"
	"sort"
)

// 组合总和 II
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var (
		freq     [][2]int
		sequence []int
		dfs      func(pos, rest int)
	)

	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func combinationSum2_2(candidates []int, target int) [][]int {
	var (
		ans       [][]int
		backtrack func(int, int)
		record    = make([]int, 0)
	)

	sort.Ints(candidates)

	backtrack = func(i int, sum int) {
		if sum > target {
			return
		} else if sum == target {
			ans = append(ans, slices.Clone(record))
			return
		}

		for j := i; j < len(candidates); j++ {
			// [i, len(candidates)) 如果有相同元素只选择第一个， 因为选择第二个的情况被选择第一个包含了
			if j == i || candidates[j] != candidates[j-1] {
				record = append(record, candidates[j])
				backtrack(j+1, sum+candidates[j])
				record = record[:len(record)-1]
			}
		}
	}
	backtrack(0, 0)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	helper.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	helper.Log(combinationSum2_2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	helper.Log(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	helper.Log(combinationSum2_2([]int{2, 5, 2, 1, 2}, 5))
}
