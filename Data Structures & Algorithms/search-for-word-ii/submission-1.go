func findWords(board [][]byte, words []string) []string {
	var res []string
	rows, cols := len(board), len(board[0])
	
	var dfs func(r, c int, i int, word string) bool

	dfs = func(r, c int, i int, word string) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[i] {
			return false
		}

		board[r][c] = '*'

		defer func() {
			board[r][c] = word[i]
		}()

		return  dfs(r+1, c, i+1, word) || dfs(r, c+1, i+1, word) || dfs(r-1, c, i+1, word) || dfs(r, c-1, i+1, word) 
	}

	for _, w := range words {
		found := false 
		for r := 0; r < rows; r ++ {
			if found {
				break
			}
			for c := 0; c < cols; c ++ {
				if board[r][c] == w[0] && dfs(r, c, 0, w) {
					res = append(res, w)
					found = true
					break
				}
			}
		}
	}

	return res
}
