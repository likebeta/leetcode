package main

import "leetcode/helper"

// 新21点
func new21Game(N int, K int, W int) float64 {
	dp := make([]float64, K+W, K+W)
	var sum float64
	for i := K; i < len(dp); i++ {
		if i <= N {
			dp[i] = 1
		}
		sum += dp[i]
	}

	for i := K - 1; i >= 0; i-- {
		dp[i] = sum / float64(W)
		sum += dp[i] - dp[i+W]
	}
	return dp[0]
}

func main() {
	helper.Log(new21Game(10, 1, 10), 1.0)
	helper.Log(new21Game(6, 1, 10), 0.6)
	helper.Log(new21Game(21, 17, 10), 0.73278)
}
