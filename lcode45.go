package main

func jump(nums []int) int {
	minJumps := 0
	farthest := 0
	current_end := 0

	for i := 0; i < len(nums); i++ {
		farthest = max(farthest, i+nums[i])
		if i > current_end {
			minJumps++
			current_end = farthest
		}
	}

	return minJumps
}
