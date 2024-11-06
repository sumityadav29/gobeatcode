package main

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	memo := make([][]int, len(matrix))

	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(matrix[i]))
	}

	longest := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			longest = max(longest, longestPath(matrix, memo, i, j, -1))
		}
	}
	return longest
}

func longestPath(matrix, memo [][]int, i, j, prevVal int) int {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) || matrix[i][j] <= prevVal {
		return 0
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	curr := matrix[i][j]

	res := 1
	res = max(res, 1+longestPath(matrix, memo, i+1, j, curr))
	res = max(res, 1+longestPath(matrix, memo, i-1, j, curr))
	res = max(res, 1+longestPath(matrix, memo, i, j+1, curr))
	res = max(res, 1+longestPath(matrix, memo, i, j-1, curr))

	memo[i][j] = res

	return res
}
