func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := make(map[int][]int)
    
    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }

    for _, pre := range prerequisites {
        preMap[pre[0]] = append(preMap[pre[0]], pre[1])
    }

    res := []int{}
    visited := make([]int, numCourses)  // 0=unvisited, 1=in-path, 2=safe

    var dfs func(course int) bool
    dfs = func(course int) bool {
        if visited[course] == 1 {
            return true  // Cycle detected
        }
        if visited[course] == 2 {
            return false  // Already processed safely
        }

        visited[course] = 1  // Mark as in current path

        for _, c := range preMap[course] {
            if dfs(c) {
                return true  // Cycle found
            }
        }

        visited[course] = 2  // Mark as safe/done
        res = append(res, course)
        return false
    }

    for i := 0; i < numCourses; i++ {
        if visited[i] == 0 {  // Only process unvisited
            if dfs(i) {
                return []int{}  // Cycle detected
            }
        }
    }
    
    return res
}