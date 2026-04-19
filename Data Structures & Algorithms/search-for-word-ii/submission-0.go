func findWords(board [][]byte, words []string) []string {
	var res []string
	rows, cold := len(board), len(board[0])
	
	var dfs func(r, c int, i int, word string) bool

	dfs = func(r, c int, i int, word string) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r > rows || c > cols || board[r][c] != word[i] {
			return false
		}

		board[r][c] = '*'

		defer func() {
			board[i][j] = word[i]
		}()

		return  dfs(r+1, c, i+1, word) || dfs(r+1, c, i+1, word) || dfs(r+1, c, i+1, word) 
	}
    
}
