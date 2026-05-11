func numIslands(grid [][]byte) int {
    res := 0
	directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) 
	dfs = func(r , c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return 
		}

		grid[r][c] = '0'
		for _, dir := range directions {
			row, col := dir[0] + r, dir[1] + c
			dfs(row, col)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '1' {
				continue
			}
			dfs(r, c)
			res++
		}
	}

	return res
}
