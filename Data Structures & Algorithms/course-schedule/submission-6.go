func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)

    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }

    for _, pre := range prerequisites {
        preMap[pre[0]] = append(preMap[pre[0]], pre[1])
    }

    visited := make([]int, numCourses)
    var dfs func(c int) bool 
    dfs = func(c int) bool {
        if len(preMap[c]) == 0 {
            return false
        }

        if visited[c] == 1 {
            return true
        }

        visited[c] = 1

        for _, course := range preMap[c] {
            if dfs(course) {
                return true
            }
        }

        visited[c] = 0
        preMap[c] = []int{}

        return false
    }

    for c := 0; c < numCourses; c++ {
        if len(preMap[c]) == 0 {
            continue
        }

        if dfs(c) {
            return false
        }
    }

    return true
}
