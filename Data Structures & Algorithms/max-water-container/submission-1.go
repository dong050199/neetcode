func maxArea(heights []int) int {
	left , right := 0 , len(heights) - 1
	res := 0
	for left < right {
		s := min(heights[left], heights[right]) * (right - left)
		res = max(res, s)

		if heights[left] > heights[right] {
			right--
		}else {
			left++
		}
	}
	
	return res
}
