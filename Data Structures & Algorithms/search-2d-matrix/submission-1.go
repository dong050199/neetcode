func searchMatrix(matrix [][]int, target int) bool {
	arr := []int{}

	for _, ma := range matrix {
		arr = append(arr, ma...)
	}

	// then do binary search
	lo, hi := 0, len(arr)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] > target {
			hi = mid -1
		} else {
			lo = mid + 1
		}
	}
	return false
}
