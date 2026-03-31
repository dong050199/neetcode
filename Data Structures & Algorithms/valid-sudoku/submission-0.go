func isValidSudoku(board [][]byte) bool {
// o(n2)
    cols := make([]map[byte]bool,9)
    rows := make([]map[byte]bool,9)
    boxes := make([]map[byte]bool,9)
        
    for i := 0 ; i < 9 ; i ++ {
        cols[i] = make(map[byte]bool)
        rows[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0 ; r < 9 ; r ++ {
        for c := 0; c < 9; c ++ {

            val := board[r][c]
            if val == '.' {
                continue
            }

            boxIndex := (r/3) * 3 + c/3

            if rows[r][val] || cols[c][val] || boxes[boxIndex][val] {
                return false
            }

            rows[r][val] = true
            cols[c][val]  = true
            boxes[boxIndex][val] = true
        }
    }

    return true
}
