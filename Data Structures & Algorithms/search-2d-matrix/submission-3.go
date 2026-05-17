func searchMatrix(matrix [][]int, target int) bool {
	// first we need to reduce the search space 
	rows, cols := len(matrix), len(matrix[0])
	selectedRow := []int{}
	for r := 0; r < rows; r++ {
		if matrix[r][0] <= target && target <= matrix[r][cols-1] {
			selectedRow = matrix[r]
		}
	}

	if len(selectedRow) == 0 {
		return false
	}

	l, r := 0, cols - 1
	for l <= r {
		mid := l + (r-l)/2
		if selectedRow[mid] == target {
			return true
		}

		if selectedRow[mid] > target {
			r = mid -1
		} else {
			l = mid + 1
		}
	}

	return false
}
