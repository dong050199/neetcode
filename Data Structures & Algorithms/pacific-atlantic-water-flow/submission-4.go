func pacificAtlantic(heights [][]int) [][]int {
    // need to get all pacific and alatis 
	rows, cols := len(heights), len(heights[0])
	directions := [][]int{{0, 1},{0,-1},{1,0},{-1,0}}

	alt := make([][]bool,rows)
	pac := make([][]bool,rows)

	for i := range alt {
		alt[i] = make([]bool, cols)
		pac[i] = make([]bool, cols)
	}

	altQ := [][]int{}
	pacQ := [][]int{}

	for r := 0; r < rows; r++ {
		alt[r][cols-1] = true
		pac[r][0] = true
		altQ = append(altQ, []int{r, cols - 1})
		pacQ = append(pacQ, []int{r, 0})
	}

	for c := 0; c < cols; c++ {
		alt[rows-1][c] = true
		pac[0][c] = true
		altQ = append(altQ, []int{rows-1, c})
		pacQ = append(pacQ, []int{0, c})
	}

	var bfs func(cache [][]bool, queue [][]int)
	bfs = func(cache [][]bool, queue [][]int) {
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				row, col := cur[0] + dir[0], cur[1] + dir[1]
				if row >= 0 && col >= 0 && row < rows && col < cols && !cache[row][col] && heights[row][col] >= heights[cur[0]][cur[1]]  {
					queue = append(queue, []int{row, col})
					cache[row][col] = true
				}
			}
		}
	}

	bfs(alt, altQ)
	bfs(pac, pacQ)

	res := [][]int{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if alt[r][c] == true && pac[r][c] == true {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
