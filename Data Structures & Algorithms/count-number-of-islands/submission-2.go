func numIslands(grid [][]byte) int {
	// DS graph
	// this DFS function will check if an land in grid is belong to an island or not

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	
	var dfs func(i, j int) 
	dfs = func(i, j int)  {
		if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0' // using the origin grid as the cache

		for _, dir := range directions {
			dfs(i + dir[0], j + dir[1])
		}
	}

	res := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(i, j) // update cache
				res++
			}
		}
	}

	return res
}
