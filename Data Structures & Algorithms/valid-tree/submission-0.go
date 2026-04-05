func validTree(n int, edges [][]int) bool {
	
	mapNode := make(map[int][]int)
	for i := 0; i < n; i++ {
		mapNode[i] = []int{}
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		mapNode[a] = append(mapNode[a], b)
		mapNode[b] = append(mapNode[b], a)
	}

	var bfs func(n int, parent int) bool

	visited := make(map[int]bool)

	bfs = func(n int, parent int) bool {
		if visited[n] {
			return false
		}

		visited[n] = true
		
		for _, node := range mapNode[n] {
			if node == parent {
				continue
			}

			if !bfs(node, n) {
				return false
			}
		}

		return true
	}

	return bfs(0, -1) && len(visited) == n
}
