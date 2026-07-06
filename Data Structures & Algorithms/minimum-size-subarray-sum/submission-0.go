func minSubArrayLen(target int, nums []int) int {
	res := 0
	l := 0
	curSum := 0

	for r := 0; r < len(nums); r++ {
		curSum += nums[r]

		for curSum >= target {
			if res == 0 || res >= r-l+1 {
				res = r - l + 1
			}
			curSum -= nums[l]
			l++
		}
	}

	return res
}