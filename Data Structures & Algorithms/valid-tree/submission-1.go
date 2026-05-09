func validTree(n int, edges [][]int) bool {
    mpNode := make(map[int][]int)

	for i := 0; i < n; i++ {
		mpNode[i] = []int{}
	}

	for _, edge := range edges {
		mpNode[edge[0]] = append(mpNode[edge[0]], edge[1])
		mpNode[edge[1]] = append(mpNode[edge[1]], edge[0])
	}

	visited := make([]int, n)
	var dfs func(node, parent int) bool
	dfs = func(node , parent int) bool {
		if visited[node] == 1 {
			return true
		}

		if visited[node] == 2 {
			return false
		}

		visited[node] = 1

		for _, n := range mpNode[node] {
			if n == parent {
				continue
			}
			if dfs(n, node) {
				return true
			}
		}


		visited[node] = 2
		return false
	}

	for i := 0; i < n; i++ {
		if visited[i] != 0 {
			continue
		}
		if dfs(i, -1) {
			return false
		}
	}

	return true
}
