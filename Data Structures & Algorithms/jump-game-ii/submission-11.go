func jump(nums []int) int {
    l, r := 0, 0
	res := 0

	for r < len(nums) - 1 {
		curMax := 0
		for i := l; i <= r; i++ {
			curMax = max(curMax, nums[i] + i)
		}
		l = r + 1
		r = curMax
		res++
	}

	return res
}
