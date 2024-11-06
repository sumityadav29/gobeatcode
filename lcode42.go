package main

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxLeft := 0
	maxRight := 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				water = water + maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				water = water + maxRight - height[right]
			}
			right--
		}
	}

	return water
}
