func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}

	for r := 1; r < rows; r++ {
		if obstacleGrid[r][0] != 1 {
			dp[r][0] = dp[r-1][0]
		} else {
			break
		}
	}

	for c := 1; c < cols; c++ {
		if obstacleGrid[0][c] != 1 {
			dp[0][c] = dp[0][c-1]
		} else {
			break
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if obstacleGrid[r][c] == 1 {
				continue
			}
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[rows-1][cols-1]
}