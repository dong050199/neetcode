func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])
	directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

	var dfs func(r, c int) 
	dfs = func(r, c int) {
		if r < 1 || c < 1 || r >= rows - 1 || c >= cols - 1 || board[r][c] != 'O' {
			return 
		}  

		board[r][c] = '*'
		for _, dir := range directions {
			row, col := dir[0] + r, dir[1] + c
			dfs(row, col)
		}
	}

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if board[r][c] == 'O' {
				dfs(r, c)
			}
		}
	}


	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if board[r][c] == '*' {
				board[r][c] = 'X'
			}
		}
	}
}
