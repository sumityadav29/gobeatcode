package main

func rob(nums []int) int {
	last := 0
	secondLast := 0

	for _, v := range nums {
		last, secondLast = max(v+secondLast, last), last
	}

	return max(last, secondLast)
}
