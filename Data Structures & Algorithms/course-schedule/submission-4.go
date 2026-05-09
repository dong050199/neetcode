func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make(map[int][]int)

	for i := range numCourses {
		courses[i] = []int{}
	}

	for _, c := range prerequisites {
		courses[c[0]] = append(courses[c[0]], c[1])
	}

	visited := make(map[int]bool)

	var dfs func(c int) bool
	dfs = func(course int) bool {
		if len(courses[course]) == 0 {
			return true
		}

		if visited[course] {
			return false
		}

		visited[course] = true

		for _, c := range courses[course] {
			dfs(c)
		}

		return false
	}

	for c := 0; c < numCourses; c++ {
		if dfs(c) {
			return true
		}
	}

	return false
}
