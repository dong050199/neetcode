func solve(board [][]byte) {
	directions := [][]int{{1, 0},{0, 1},{-1, 0},{0, -1}}
	var dfs func(i, j int) 
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
			return
		} 

		board[i][j] = '*'
		
		for _, dir := range directions {
			r , c := i + dir[0] , j + dir[1]
			dfs(r, c)
		}
	}

	for i := range board {
		for j := range board[i] {
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[0]) - 1{
				if board[i][j] == 'O' {
					dfs(i, j)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}
