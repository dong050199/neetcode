func largestRectangleArea(heights []int) int {
    maxArea := 0
    
    for i := 0; i < len(heights); i++ {
        minHeight := heights[i]
        
        // Expand from current position
        for j := i; j < len(heights); j++ {
            minHeight = min(minHeight, heights[j])
            width := j - i + 1
            area := minHeight * width
            maxArea = max(maxArea, area)
        }
    }
    
    return maxArea
}