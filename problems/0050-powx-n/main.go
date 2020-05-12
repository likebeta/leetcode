package main

import (
    "leetcode/helper"
)

// x 的 n 次幂
func myPow(x float64, n int) float64 {
    if n >= 0 {
        return recursion(x, n)
    }
    return 1.0 / recursion(x, -n)
}

func recursion(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }

    half := recursion(x, n/2)
    if n%2 == 0 {
        return half * half
    }
    return half * half * x
}

func myPow2(x float64, n int) float64 {
    if n >= 0 {
        return process(x, n)
    }
    return 1.0 / process(x, -n)
}

func process(x float64, n int) float64 {
    ans := 1.0
    pow := x
    for n > 0 {
        if (n-1)&1 == 0 {
            ans *= pow
        }
        pow *= pow
        n = n >> 1
    }
    return ans
}

func main() {
    helper.Log(myPow(2.0, 10), myPow2(2.0, 10))
    helper.Log(myPow(2.1, 3), myPow2(2.1, 3))
    helper.Log(myPow(2.0, -2), myPow2(2.0, -2))
}
