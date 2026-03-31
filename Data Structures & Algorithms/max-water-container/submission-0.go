func maxArea(heights []int) int {
    // using two pointer
    l := 0
    r := len(heights) - 1
    max := 0

    for r > l {
        w := (r - l) * min(heights[l], heights[r])
        if w > max {
            max = w
        }

        switch {
            case heights[l] > heights[r]:
                r--
            case heights[l] < heights[r]:
                l++
            default:
                l++
        }

    }
    return max
}
