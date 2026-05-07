func orangesRotting(grid [][]int) int {
	queue := [][]int{}

	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r ++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	res := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		newQ := [][]int{}
		for _, cell := range queue {
			for _, dir := range directions {
				r, c := cell[0] + dir[0], cell[1] + dir[1]
				if r >= 0 && c >= 0 && r < rows && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					newQ = append(newQ, []int{r, c})
				}
			}
		}

		if len(newQ) > 0 {
			res++
		}
		queue = newQ
	}

	for r := 0; r < rows; r ++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return res
}
