package main

import (
	"leetcode/helper"
	"sort"
)

// 重新安排行程
func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		ans []string
		dfs func(curr string)
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		ans = append(ans, curr)
	}

	dfs("JFK")
	if n := len(ans); n > 1 {
		for i := 0; i < n/2; i++ {
			ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
		}
	}
	return ans
}

func main() {
	// ["JFK", "MUC", "LHR", "SFO", "SJC"]
	helper.Log(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	// ["JFK","ATL","JFK","SFO","ATL","SFO"]
	helper.Log(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
