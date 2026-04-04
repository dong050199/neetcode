func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int, numCourses)
	for c := 0 ; c < numCourses; c ++ {
		preMap[c] = []int{}
	}

	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	output := []int{}
	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if visiting[course] {
			return false
		}

		if visited[course] {
			return true
		}

		visiting[course] = true

		for _,c := range preMap[course] {
			if !dfs(c) {
				return false
			}
		}

		visiting[course] = false
		visited[course] = true
		output = append(output, course)
		return true
	}

	for c := 0 ; c < numCourses ; c++ {
		if !dfs(c) {
			return []int{}
		}
	}

	return output
}
