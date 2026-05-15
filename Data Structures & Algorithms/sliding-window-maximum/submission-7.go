func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	q := make([]int, len(nums))
	q[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		q[i] = max(q[i-1], nums[i])		
	}

	res := []int{}
	for r := k - 1; r < len(nums); r++ {
		res = append(res, q[r])
	}

    return res
}

// [1,2,1,0,4,2,6]
// [1,2,2,2,4,4,6]
// [1] -> 2
// [2] -> 2 
// [3] -> 4
// 
