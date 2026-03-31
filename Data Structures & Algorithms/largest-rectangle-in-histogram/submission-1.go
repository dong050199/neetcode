func largestRectangleArea(heights []int) int {
    var maxArea int
    for i := 0; i < len(heights); i ++ {
        minHeight := heights[i]
        for j := i ; j < len(heights); j ++ {
            minHeight = min(minHeight, heights[j])
            area := (j - i + 1) * minHeight
            maxArea = max(maxArea, area)
        }
    }
    return maxArea
}

// how to solve this problem
// brute force solution:
// first we need to return area = max heigh * max width
// the area for each run is max width * min heigh
// first index: 1 * 7 = 7 -> 7
// the second index: 1 * 2 = 2 < 7 -> 7
// 1 * 3 = 3 < 7 -> 7
// 1 * 4 ..... 1*6 -> 7

// second postion: 1 * 1 = 1
// .... 1 * 5 -> 7
// third postion: 7 * 1 = 7
// 2 * 2 * 2 * 2 = 8 -> max = 7