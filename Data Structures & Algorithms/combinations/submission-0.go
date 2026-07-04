func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	}

	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	selected := make([]int, n)
	res := [][]int{}
	var dfs func(idx int, combination []int, numberLeft int)
	dfs = func(idx int, combination []int, numberLeft int) {
		if numberLeft == 0 {
			tmp := make([]int, len(combination))
			copy(tmp, combination)
			res = append(res, tmp)
			return
		}

		if idx >= len(arr) {
			return
		}

		selected[idx] = 1
		dfs(idx+1, append(combination, arr[idx]), numberLeft-1)
		selected[idx] = 0
		dfs(idx+1, combination, numberLeft)
	}

	dfs(0, []int{}, k)

	return res
}