func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }

    for _, pre := range prerequisites {
        preMap[pre[0]] = append(preMap[pre[0]], pre[1])
    }

    var dfs func(course int) bool
    visited := make([]int, numCourses)
    dfs = func(course int) bool {
        if visited[course] == 1 {
            return true
        }

        if visited[course] == 2 {
            return false
        }

        visited[course] = 1 
        for _, c := range preMap[course] {
            if dfs(c) {
                return true
            }
        }

        visited[course] = 2
        return false
    }

    for i := 0; i < numCourses; i++ {
        if visited[i] > 0 {
            continue
        }
        if dfs(i) {
            return false
        }
    }

    return true
}
