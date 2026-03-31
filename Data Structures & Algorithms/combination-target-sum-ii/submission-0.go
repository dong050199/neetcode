func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)  
	dfs(candidates, 0, target, []int{}, &res)
	return res
}

func dfs(candidates []int, i int, target int, subset []int, res *[][]int) {
	total := 0
	for _, s := range subset {
		total += s
	}

	if target == total {
		tmp := make([]int, len(subset))
		copy(tmp, subset)
		*res = append(*res, tmp)
		return
	}

	if i >= len(candidates) || total > target { 
		return
	}

	subset = append(subset, candidates[i])
	dfs(candidates, i+1, target, subset, res) 

	subset = subset[:len(subset)-1]
	
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	
	dfs(candidates, i+1, target, subset, res)
}