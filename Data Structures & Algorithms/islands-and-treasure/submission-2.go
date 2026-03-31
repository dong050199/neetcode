func islandsAndTreasure(grid [][]int) {
	// the main idea is using BFS, each way we increse the adjance by 1
	queue := [][]int{}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if len(queue) == 0 {
		return 
	}

	for len(queue) > 0 {
		cur := queue[0] 
		queue = queue[1:]

		row , col := cur[0], cur[1]

		for _, dir := range directions {
			r, c := row + dir[0] , col + dir[1]
			if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 2147483647 {
				continue
			} 
			queue = append(queue, []int{r, c})
			grid[r][c] = grid[row][col] + 1
		}
	}
}
