package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(dp)-1] = 1

	for i := len(dp) - 2; i >= 0; i-- {
		maxLISLengthForI := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				maxLISLengthForI = max(maxLISLengthForI, dp[j])
			}
		}
		dp[i] = 1 + maxLISLengthForI
	}

	return dp[0]
}
