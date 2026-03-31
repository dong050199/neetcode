func rotate(matrix [][]int)  {
	// rotate
	for i , j := 0, len(matrix) - 1; i < j; i, j = i + 1, j - 1 {
		matrix[i] , matrix[j] = matrix[j], matrix[i]
	}
	// transpose
	for i := 0; i < len(matrix); i ++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	} 
}
