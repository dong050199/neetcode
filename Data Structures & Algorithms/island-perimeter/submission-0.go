func islandPerimeter(grid [][]int) int {
	// this one can use dfs recursion
	res := 0
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				continue
			}

			for _, dir := range directions {
				row, col := dir[0]+r, dir[1]+c
				if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == 0 {
					res++
				}
			}
		}
	}

	return res
}