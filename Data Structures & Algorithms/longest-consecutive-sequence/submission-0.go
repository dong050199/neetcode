func longestConsecutive(nums []int) int {
    max := 0
    m := make(map[int]int)

    for _, num := range nums {
        if m[num] == 0 {
            

        left := m[num-1]
        right := m[num+1]

        sum := left + right + 1
        m[num] = sum
        m[num-left] = sum
        m[num+right] = sum


        if sum > max {
            max = sum
        }
        }
    }
    return max
}
