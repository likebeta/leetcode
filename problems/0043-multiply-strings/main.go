package main

import (
	"leetcode/helper"
)

// 字符串相乘
var mapChar []byte
var mapByte []int

func init() {
	mapChar = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	mapByte = make([]int, '9'+1)
	mapByte['0'] = 0
	mapByte['1'] = 1
	mapByte['2'] = 2
	mapByte['3'] = 3
	mapByte['4'] = 4
	mapByte['5'] = 5
	mapByte['6'] = 6
	mapByte['7'] = 7
	mapByte['8'] = 8
	mapByte['9'] = 9
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1, len2 := len(num1), len(num2)
	length := len1 + len2
	arr := make([]int, length)
	for i := len1 - 1; i >= 0; i-- {
		if num1[i] != '0' {
			b := mapByte[num1[i]]
			for j := len2 - 1; j >= 0; j-- {
				if num2[j] != '0' {
					mul := b * mapByte[num2[j]]
					arr[i+j+1] += mul
				}
			}
		}
	}

	var sum, carry int
	ans := make([]byte, length)
	for length > 0 {
		length--
		sum = arr[length] + carry
		if sum > 9 {
			ans[length] = mapChar[sum%10]
			carry = sum / 10
		} else {
			ans[length] = mapChar[sum]
			carry = 0
		}
	}
	if ans[0] == '0' {
		return string(ans[1:])
	} else {
		return string(ans)
	}
}
func main() {
	helper.Assert(multiply("2", "3") == "6")
	helper.Assert(multiply("123", "456") == "56088")
}
