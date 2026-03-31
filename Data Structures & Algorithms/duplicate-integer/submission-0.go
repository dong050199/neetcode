func hasDuplicate(nums []int) bool {
    m := make(map[int]int)

    for _, n := range nums {
        m[n]++
    }

    for _,v := range m {
        if v > 1 {
            return true
        }
    }

    return false
}
