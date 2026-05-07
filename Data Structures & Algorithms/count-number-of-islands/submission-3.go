func numIslands(grid [][]byte) int {
    directions := [][]int{{1, 0},{-1, 0},{0, 1},{0, -1}}
	rows, cols := len(grid), len(grid[0])
	res := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || c >= cols || r >= rows || grid[r][c] != '1' {
			return 1
		}

		grid[r][c] = '0'

		for _, dir := range directions {
			row := dir[0] + r
			col := dir[1] + c
			dfs(row, col)
		}

		return 1
	}


	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '0' {
				continue
			}
			res += dfs(r, c)
		}
	}

	return res
}
