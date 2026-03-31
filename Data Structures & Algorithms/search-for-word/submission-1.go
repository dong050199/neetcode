func exist(board [][]byte, word string) bool {
	for r := range board {
		for c := range board[0] {
			if dfs(board, word, r, c , 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, r, c, i int) bool {
	if i == len(word) {
		return true
	}

	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}

	if board[r][c] != word[i] {
		return false
	}

	tmp := board[r][c]
	board[r][c] = '*'

	found := dfs(board, word, r+1, c , i+1) || dfs(board, word, r, c+1 , i+1) || dfs(board, word, r-1, c , i+1) || dfs(board, word, r, c-1 , i+1)
	
	board[r][c] = tmp
	return found
}


