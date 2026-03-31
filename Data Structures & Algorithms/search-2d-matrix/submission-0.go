func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])
    // single pass
    l , r := 0, rows * cols - 1
    for l <= r {
        m := l + (r - l)/2
        // re calculate row and col of median
        row, col := m / cols, m % cols
        if matrix[row][col] == target {
            return true
        }
        if matrix[row][col] > target {
            r = m - 1
            continue
        }
        l = m + 1
    }
    return false
}