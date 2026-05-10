func searchMatrix(matrix [][]int, target int) bool {
	// convert than search?
	rows, cols := len(matrix), len(matrix[0])

	// do binary seach for rows first then do for row
	selectedRow := -1
	for r := 0; r < rows; r++ {
		if matrix[r][0] <= target && matrix[r][cols-1] >= target {
			selectedRow = r
		}
	}

	if selectedRow == -1 {
		return false
	}

	l, r := 0, cols - 1
	cur := matrix[selectedRow]
	for l <= r {
		mid := l + (r - l)/2 
		if cur[mid] == target {
			return true
		}

		if cur[mid] > target {
			r = mid -1
		} else {
			l = mid + 1
		}
	}

	return false
}
