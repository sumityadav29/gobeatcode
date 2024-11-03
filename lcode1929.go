package main

func getConcatenation(nums []int) []int {
	length := 2 * len(nums)
	concatenatedSlice := make([]int, length, length)

	for i := 0; i < length; i++ {
		if i < len(nums) {
			concatenatedSlice[i] = nums[i]
		} else {
			concatenatedSlice[i] = nums[i-len(nums)]
		}
	}

	return concatenatedSlice
}
