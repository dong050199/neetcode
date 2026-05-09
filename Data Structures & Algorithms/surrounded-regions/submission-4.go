func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	directions := [][]int{{0,1},{0,-1},{-1,0},{1,0}}

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
			return 
		} 

		board[r][c] = '*'

		for _, dir := range directions {
			row, col := dir[0] + r, dir[1] + c 
			dfs(row, col)
		}
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0)
		dfs(r, cols-1)
	}

	for c := 0; c < cols; c++ {
		dfs(0, c)
		dfs(rows-1, c)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}

			if board[r][c] == '*' {
				board[r][c] = 'O'
			}
		}
	}
}
