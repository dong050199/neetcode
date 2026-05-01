func exist(board [][]byte, word string) bool {
	directions := [][]int{{1, 0},{-1, 0},{0, 1},{0, -1}}
	rows, cols := len(board), len(board[0])
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[idx] {
			return false
		}

		if idx == len(word) - 1 {
			return true
		}

		for _, dir := range directions {
			row, col := r + dir[0], c + dir[1]
			if dfs(row, col, idx + 1) {
				return true
			}
		}

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
