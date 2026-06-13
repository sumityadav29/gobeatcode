package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			for k := 0; k < len(strs); k++ {
				zeroes := strings.Count(strs[k], "0")
				ones := len(strs[k]) - zeroes

				if i-zeroes >= 0 && j-ones >= 0 {
					dp[i][j] = max(dp[i][j], 1+dp[i-zeroes][j-ones])
				}

			}
		}
	}

	return dp[m][n]

}
