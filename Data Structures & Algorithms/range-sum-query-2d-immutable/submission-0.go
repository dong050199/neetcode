

type NumMatrix struct {
    ps [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    ps := make([][]int, m + 1)
    for i := range ps {
        ps[i] = make([]int, n + 1)
    }

    for r := 1; r <= m; r++ {
        for c := 1; c <= n; c++ {
            ps[r][c] = matrix[r-1][c-1] + ps[r][c-1] + ps[r-1][c] - ps[r-1][c-1]
        }
    }

    return NumMatrix{
        ps: ps,
    }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.ps[row2 + 1][col2 + 1] - this.ps[row1][col2 + 1] - this.ps[row2 + 1][col1] + this.ps[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */