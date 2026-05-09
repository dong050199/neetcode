func countComponents(n int, edges [][]int) int {
	mpNode := make(map[int][]int)

	for i := 0; i < n; i++ {
		mpNode[i] = []int{}
	}

	for _, edge := range edges {
		mpNode[edge[0]] = append(mpNode[edge[0]], edge[1])
		mpNode[edge[1]] = append(mpNode[edge[1]], edge[0])
	}

	visited := make([]int, n)
	var dfs func(node, parent int) 

	dfs = func(node, parent int) {
		if visited[node] != 0 {
			return 
		}

		visited[node] = 1

		for _, n := range mpNode[node] {
			if n == parent {
				continue
			}
			dfs(n, node)
		}

		visited[node] = 2
	} 

	res := 0
	for i := 0; i < n; i++ {
		if visited[i] > 0 {
			continue
		}
		dfs(i, -1)
		res++
	}
	return res
}
