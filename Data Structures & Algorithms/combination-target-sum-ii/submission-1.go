func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates) 
	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		
		if i >= len(candidates) || total > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			if total + candidates[j] > target {
				break
			}

			cur = append(cur, candidates[j])
			dfs(j + 1, cur, total + candidates[j])  // ← j+1, not i+1
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, []int{}, 0)
	return res
}