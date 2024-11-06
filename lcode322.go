package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
