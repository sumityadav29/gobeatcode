package main

func uniquePathsWithObstacles(grid [][]int) int {
	paths := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		paths[i] = make([]int, len(grid[i]))
		for j := 0; j < len(paths[i]); j++ {
			if grid[i][j] == 1 {
				paths[i][j] = 0
			} else if i == 0 && j > 0 {
				paths[i][j] = paths[i][j-1]
			} else if j == 0 && i > 0 {
				paths[i][j] = paths[i-1][j]
			} else if i == 0 || j == 0 {
				paths[i][j] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}

	return paths[len(paths)-1][len(paths[0])-1]
}
