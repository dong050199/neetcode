func maxAreaOfIsland(grid [][]int) int {
	// same with the number of island

	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	var dfs func (i, j int) int 

	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
			return 0
		}
		
		grid[i][j] = 0

		res := 1

		for _, dir := range directions {
			res += dfs(i + dir[0], j + dir[1])
		}

		return res
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}

	return res
}
