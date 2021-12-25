package main

import (
	"leetcode/helper"
)

// 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}

func testOne(gasIn, costIn string, ans int) {
	gas := helper.ParseArray(gasIn)
	cost := helper.ParseArray(costIn)
	helper.Assert(canCompleteCircuit(gas, cost) == ans)
}

func main() {
	testOne("[1,2,3,4,5]", "[3,4,5,1,2]", 3)
	testOne("[2,3,4]", "[3,4,3]", -1)
}
