func islandsAndTreasure(grid [][]int) {
	inf := 2147483647
	treasure := 0
	q := [][]int{}
	rows, cols := len(grid), len(grid[0])
	
	// Find all treasure cells and add to queue
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == treasure {
				q = append(q, []int{r, c})
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	
	for len(q) > 0 {
		newq := [][]int{}
		for _, cell := range q {
			r, c := cell[0], cell[1]
			for _, dir := range directions {
				row, col := dir[0] + r, dir[1] + c
				if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == inf {
					grid[row][col] = grid[r][c] + 1 
					newq = append(newq, []int{row, col})
				}
			}
		}
		
		q = newq  // Move to next level
	}
}