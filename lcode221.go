package main

func maximalSquare(matrix [][]byte) int {

	dp := make([][]int, len(matrix)+1)
	maxSide := 0

	for i := 0; i <= len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
