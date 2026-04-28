func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	visited := make(map[[2]int]bool)

	directions := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || 
		   visited[[2]int{r, c}] || board[r][c] != word[i] {
			return false
		}

		visited[[2]int{r, c}] = true

		for _, dir := range directions {
			row := r + dir[0]
			col := c + dir[1]
			if dfs(row, col, i+1) {
				visited[[2]int{r, c}] = false  
				return true
			}
		}

		visited[[2]int{r, c}] = false
		return false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}