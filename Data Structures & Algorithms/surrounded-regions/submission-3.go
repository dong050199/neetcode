func solve(board [][]byte) {
    if len(board) == 0 {
        return
    }
    
    rows, cols := len(board), len(board[0])
    directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

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

    // DFS from all border 'O' cells
    // Top and bottom rows
    for c := 0; c < cols; c++ {
        if board[0][c] == 'O' {
            dfs(0, c)
        }
        if board[rows-1][c] == 'O' {
            dfs(rows-1, c)
        }
    }

    // Left and right columns
    for r := 0; r < rows; r++ {
        if board[r][0] == 'O' {
            dfs(r, 0)
        }
        if board[r][cols-1] == 'O' {
            dfs(r, cols-1)
        }
    }

    // Convert: '*' stays 'O', 'O' becomes 'X'
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if board[r][c] == '*' {
                board[r][c] = 'O'  // Safe cells stay O
            } else if board[r][c] == 'O' {
                board[r][c] = 'X'  // Surrounded cells become X
            }
        }
    }
}