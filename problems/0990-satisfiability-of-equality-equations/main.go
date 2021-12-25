package main

import (
	"leetcode/helper"
)

// 等式方程的可满足性
func equationsPossible(equations []string) bool {
	leaders := make([]byte, int('z'+1))
	for i := byte('a'); i < 'z'; i++ {
		leaders[i] = i
	}

	for _, str := range equations {
		if str[1] == '=' {
			union(leaders, str[0], str[3])
		}
	}

	for _, str := range equations {
		if str[1] == '!' {
			if find(leaders, str[0]) == find(leaders, str[3]) {
				return false
			}
		}
	}
	return true
}

func union(parent []byte, i1, i2 byte) {
	parent[find(parent, i1)] = find(parent, i2)
}

func find(parent []byte, index byte) byte {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}

func main() {
	helper.Assert(equationsPossible([]string{"a==b", "b!=a"}) == false)
	helper.Assert(equationsPossible([]string{"b==a", "a==b"}) == true)
	helper.Assert(equationsPossible([]string{"a==b", "b==c", "a==c"}) == true)
	helper.Assert(equationsPossible([]string{"a==b", "b!=c", "c==a"}) == false)
	helper.Assert(equationsPossible([]string{"c==c", "b==d", "x!=z"}) == true)
}
