package main

import (
	"leetcode/helper"
)

// 组合
func combine(n int, k int) [][]int {
	var (
		ans [][]int
		dfs func(int, []int)
	)

	dfs = func(index int, record []int) {
		if len(record) == k {
			comb := make([]int, k)
			copy(comb, record)
			ans = append(ans, comb)
			return
		}

		if index > n {
			return
		}

		dfs(index+1, record)
		record = append(record, index)
		dfs(index+1, record)
	}

	dfs(1, nil)
	return ans
}

func combine2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := make([]int, k+1)
	for i := 0; i < k; i++ {
		temp[i] = i + 1
	}
	temp[k] = n + 1

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}

func main() {
	helper.Log(combine(4, 2))
	helper.Log(combine2(4, 2))
	helper.Log(combine(5, 4))
	helper.Log(combine2(5, 4))
}
