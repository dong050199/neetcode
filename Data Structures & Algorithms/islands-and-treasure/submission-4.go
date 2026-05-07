func islandsAndTreasure(grid [][]int) {
	const (
		inf     = 2147483647
		treasure = 0
	)

	queue := [][]int{}
	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == treasure {
				queue = append(queue, []int{r, c})
			}
		}
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	level := 1
	for len(queue) > 0 {
		newQ := [][]int{}
		for _, cell := range queue {
			for _, dir := range directions {
				r, c := cell[0]+dir[0], cell[1]+dir[1]
				if r >= 0 && c >= 0 && r < rows && c < cols && grid[r][c] == inf {
					grid[r][c] = level
					newQ = append(newQ, []int{r, c})
				}
			}
		}
		queue = newQ
		level++
	}
}