func numIslands(grid [][]byte) int {
    res := 0
	rows, cols := len(grid), len(grid[0])
	var dfs func (r, c int) 
	directions := [][]int{{0, 1},{0, -1},{1, 0},{-1,0}}
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
			return 
		}
		grid[r][c] = '0'
		for _, dir := range directions {
			row, col := r + dir[0], c + dir[1]
			dfs(row, col)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '0' {
				continue
			}
			dfs(r, c)
			res++
		}
	}

	return res
}
