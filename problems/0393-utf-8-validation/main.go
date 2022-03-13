package main

import "leetcode/helper"

// UTF-8 编码验证
const (
	mask1 = 1 << 7
	mask2 = 3 << 6
)

func getBytes(num int) int {
	if num&mask1 == 0 {
		return 1
	}
	n := 0
	for mask := mask1; num&mask != 0; mask >>= 1 {
		n++
	}
	if n >= 2 && n <= 4 {
		return n
	}
	return -1
}

func validUtf8(data []int) bool {
	n, m := 1, len(data)
	for i := 0; i < m; i += n {
		n = getBytes(data[i])
		if n <= 0 || i+n > m {
			return false
		}

		for _, ch := range data[i+1 : i+n] {
			if ch&mask2 != mask1 {
				return false
			}
		}
	}
	return true
}

func testOne(in string, result bool) {
	nums := helper.ParseArray(in)
	ans := validUtf8(nums)
	helper.Assert(result == ans)
}

func main() {
	testOne("[197,130,1]", true)
	testOne("[235,140,4]", false)

}
