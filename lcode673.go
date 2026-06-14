package main

func findNumberOfLIS(nums []int) int {
	gmax := 1
	gmaxCount := 1
	dp := make([]int, len(nums))
	dpcount := make([]int, len(nums))
	dp[0] = 1
	dpcount[0] = 1

	for i := 1; i < len(nums); i++ {
		maxForITillNow := 0
		countOfMaxForITillNow := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] > maxForITillNow {
					maxForITillNow = dp[j]
					countOfMaxForITillNow = dpcount[j]
				} else if dp[j] == maxForITillNow {
					countOfMaxForITillNow += dpcount[j]
				}
			}
		}
		dp[i] = 1 + maxForITillNow
		if maxForITillNow == 0 {
			dpcount[i] = 1
		} else {
			dpcount[i] = countOfMaxForITillNow
		}

		if dp[i] > gmax {
			gmax = dp[i]
			gmaxCount = dpcount[i]
		} else if dp[i] == gmax {
			gmaxCount += dpcount[i]
		}
	}

	return gmaxCount
}
