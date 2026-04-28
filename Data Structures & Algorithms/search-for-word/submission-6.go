func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	directions := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[i] {
			return false
		}

		temp := board[r][c]
        board[r][c] = '#'

		for _, dir := range directions {
			row := r + dir[0]
			col := c + dir[1]
			if dfs(row, col, i+1) {
				return true
			}
		}

		board[r][c] = temp

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