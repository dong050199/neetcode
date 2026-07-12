func findClosestElements(arr []int, k int, x int) []int {
    l, r := 0, len(arr)-1

    // find first element > x
    for l <= r {
        mid := l + (r-l)/2
        if arr[mid] > x {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    left, right := r, l

    res := make([]int, 0, k)

    for len(res) < k {
        if left >= 0 && right < len(arr) {
            if abs(arr[left], x) <= abs(arr[right], x) {
                res = append(res, arr[left])
                left--
            } else {
                res = append(res, arr[right])
                right++
            }
        } else if left >= 0 {
            res = append(res, arr[left])
            left--
        } else {
            res = append(res, arr[right])
            right++
        }
    }

    sort.Ints(res)
    return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}