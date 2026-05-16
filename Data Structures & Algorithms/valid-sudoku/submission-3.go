func isValidSudoku(board [][]byte) bool {
	rows, cols := len(board), len(board[0])

	rowsArr := make([]map[byte]bool, 9)
	colsArr := make([]map[byte]bool, 9)
	squareArr := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowsArr[i] = make(map[byte]bool)
		colsArr[i] = make(map[byte]bool)
		squareArr[i] = make(map[byte]bool)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == '.' {
				continue
			}

			if rowsArr[r][board[r][c]] {
				return false
			} 
			rowsArr[r][board[r][c]] = true

			if colsArr[c][board[r][c]] {
				return false
			}
			colsArr[c][board[r][c]] = true

			sq := (r/3)*3 + c/3
			if squareArr[sq][board[r][c]] {
				return false
			}
			squareArr[sq][board[r][c]] = true
		}
	}
	return true
}
