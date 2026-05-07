func maxAreaOfIsland(grid [][]int) int {
	directions := [][]int{{0, 1},{0, -1},{1, 0},{-1, 0}}
	res := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) int

	dfs = func(r, c int) int {
		if r >= rows || c >= cols || r < 0 || c < 0 || grid[r][c] == 0 {
			return 0
		}

		tmp := 1
		grid[r][c] = 0

		for _, dir := range directions {
			row, col := dir[0] + r, dir[1] + c
			tmp += dfs(row, col)
		}

		return tmp
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				continue
			}
			res = max(res, dfs(r, c))
		}
	}

	return res
}
