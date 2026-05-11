func countComponents(n int, edges [][]int) int {
    mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		mp[i] = []int{}
	}

	for _, edge := range edges {
		mp[edge[0]] = append(mp[edge[0]], edge[1])
		mp[edge[1]] = append(mp[edge[1]], edge[0])
	}

	visited := make([]int, n)
	var dfs func(node, parent int) 
	dfs = func(node, parent int) {
		if visited[node] == 1 {
			return 
		}

		if visited[node] == 2 {
			return
		}

		visited[node] = 1

		for _, edge := range mp[node] {
			if edge == parent {
				continue
			}
			dfs(edge, node)
		}

		visited[node] = 2
	}

	res := 0

	for i := 0; i < n; i ++ {
		if visited[i] > 0 {
			continue
		}
		dfs(i, -1)
		res++
	}
	return res
}
