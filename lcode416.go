package main

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	dp := make([][]bool, (sum/2)+1)

	for i := range dp {
		dp[i] = make([]bool, len(nums))
	}

	for i := range dp {
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = true
				continue
			}
			sumReqd := i
			isPossible := false

			if sumReqd == nums[j] {
				isPossible = isPossible || true
			}

			if sumReqd-nums[j] > 0 && j-1 >= 0 {
				isPossible = isPossible || dp[i-nums[j]][j-1]
			}

			if j-1 >= 0 {
				isPossible = isPossible || dp[i][j-1]
			}

			dp[i][j] = isPossible
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
