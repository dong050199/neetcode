func orangesRotting(grid [][]int) int {
	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	var (
		res int
		memo int
	)


	for {

		queue := [][]int{}

		cal := 0
		// we need to repeat this process
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == 2 {
					queue = append(queue, []int{i, j})
				}

				if grid[i][j] == 1 {
					cal++
				}
			}
		}

		if len(queue) == 0 {
    		if cal == 0 {
        		return 0  // No rotten, no fresh → done
    		}
    		return -1    // No rotten, but fresh exists → impossible
		}

		// still having frestfood
		if cal > 0 {
			if memo == cal {
				return -1
			}
			memo = cal
			res++
		}
		// no frest food anymore
		if cal == 0 {
			return res
		}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			row, col := cur[0], cur[1]
			for _, dir := range directions {
				r, c := row + dir[0], col + dir[1]
				if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
			}
		}
	
	}

	return res
}


