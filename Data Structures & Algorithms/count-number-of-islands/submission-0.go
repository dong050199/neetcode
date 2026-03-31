func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	islands := 0

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var bfs func(r, c int)
	bfs = func(r, c int) {
		queue := [][]int{{r, c}}
		grid[r][c] = '0'

		for len(queue) > 0 {
			// get from queue
			cur := queue[0]
			// dequeu
			queue = queue[1:]
			row, col := cur[0], cur[1]

			for _, dir := range directions {
				nr, nc := row + dir[0], col + dir[1]
				if nc < 0 || nr < 0 || nr >= rows || nc >= cols || grid[nr][nc] == '0' {
					continue
				}
				queue = append(queue, []int{nr, nc})
				// mark as 0 for the next cell skip it
				grid[nr][nc] = '0'
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols;  c++ {
			if grid[r][c] == '1' {
 				bfs(r, c)
				islands++
			}
		}
	}

    return islands
}
