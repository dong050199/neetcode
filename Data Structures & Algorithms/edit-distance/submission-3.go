func minDistance(word1 string, word2 string) int {
	memo := make(map[[2]int]int)

	var dfs func(i, j int) int 
	dfs = func(i, j int) int {
		memoKey := [2]int{i, j}

		if _, exist := memo[memoKey]; exist {
			return memo[memoKey]
		}

		if i == len(word1) {
			return len(word2)  - j
		}

		if j == len(word2) {
			return len(word1) - i
		}

		if word1[i] == word2[j] {
			memo[memoKey] = dfs(i+1 , j+1)
			return memo[memoKey]
		}
		memo[memoKey] = 1 + min(dfs(i + 1, j) , dfs(i, j+1), dfs(i + 1, j +1)) 
		return memo[memoKey]
	}


	return dfs(0, 0)
}
