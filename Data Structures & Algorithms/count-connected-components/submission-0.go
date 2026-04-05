func countComponents(n int, edges [][]int) int {
	mapNode := make(map[int][]int)

	for i := 0; i < n; i ++ {
		mapNode[i] = []int{}
	}

	for _, ed := range edges {
		mapNode[ed[0]] = append(mapNode[ed[0]], ed[1])
	}

	visited := make(map[int]bool)
	var bfs func(n int)
	bfs = func(n int) {
		if visited[n] {
			return
		}

		visited[n] = true

		for _, node := range mapNode[n] {
			bfs(node)
		}

		return
	}

	var res int

	for i := 0; i < n; i ++ {
		if visited[i] {
			continue
		}
		bfs(i)
		res++
	}

	return res
}
