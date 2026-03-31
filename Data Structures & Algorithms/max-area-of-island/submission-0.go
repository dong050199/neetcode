func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    ma := 0

    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
            return 0
        }
        grid[r][c] = 0

        area := 1
        for _, dir := range directions {
            area += dfs(r+dir[0], c+dir[1])
        }
        return area
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {
                continue
            }
            area := dfs(r, c)
            ma = max(ma, area)
        }
    }

    return ma
}