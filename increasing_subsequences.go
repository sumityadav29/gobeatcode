package main

import "fmt"

func AllIncreasingSubsequences(nums []int, curr []int, index int) {
	if index == len(nums) {
		if len(curr) > 0 {
			fmt.Println(curr)
		}
		return
	}

	if len(curr) == 0 || nums[index] > curr[len(curr)-1] {
		next := make([]int, len(curr), len(curr)+1)
		copy(next, curr)
		next = append(next, nums[index])
		AllIncreasingSubsequences(nums, next, index+1)
	}

	AllIncreasingSubsequences(nums, curr, index+1)
}

func GetAllIncreasingSubsequences(nums []int, curr []int, index int, res [][]int) {
	if index == len(nums) {
		if len(curr) > 0 {
			res = append(res, curr)
		}
		return
	}

	if len(curr) == 0 || nums[index] > curr[len(curr)-1] {
		next := make([]int, len(curr), len(curr)+1)
		copy(next, curr)
		next = append(next, nums[index])
		AllIncreasingSubsequences(nums, next, index+1)
	}

	AllIncreasingSubsequences(nums, curr, index+1)
}

func GetLongestIncreasingSubsequences(nums []int, curr []int, index int) []int {
	if index == len(nums) {
		if len(curr) > 0 {
			return curr
		}
		return []int{}
	}

	var res []int

	if len(curr) == 0 || nums[index] > curr[len(curr)-1] {
		next := make([]int, len(curr), len(curr)+1)
		copy(next, curr)
		next = append(next, nums[index])
		res = GetLongestIncreasingSubsequences(nums, next, index+1)
	}

	excl := GetLongestIncreasingSubsequences(nums, curr, index+1)

	if len(excl) > len(res) {
		res = excl
	}

	return res
}
