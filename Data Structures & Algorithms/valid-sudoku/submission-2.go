func isValidSudoku(board [][]byte) bool {
	rows, cols := len(board), len(board[0])

	rowsArr 	:= make([]map[byte]bool,rows)
	colsArr 	:= make([]map[byte]bool,rows)
	squareArr 	:= make([]map[byte]bool,rows)

	for i := range board {
		rowsArr[i] 	= make(map[byte]bool)
		colsArr[i] 	= make(map[byte]bool)
		squareArr[i] = make(map[byte]bool)
	}
	
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cur := board[r][c]

			if cur == '.' {
				continue
			}

			sqr := (r/3)*3 + c/3
			// check for row, col, circle
			if rowsArr[r][cur] || colsArr[c][cur] || squareArr[sqr][cur] {
				return false
			}
			rowsArr[r][cur] 	= true
			colsArr[c][cur] 	= true
			squareArr[sqr][cur] = true
		}
	}

	return true
	
}
