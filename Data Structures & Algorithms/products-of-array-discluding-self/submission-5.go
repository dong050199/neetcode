func productExceptSelf(nums []int) []int {
	res := []int{}
	pref := make(map[int]int)
	suff := make(map[int]int)

	pref[0] = 1
	for i := 1; i < len(nums); i++ {
		pref[i] = nums[i-1] * pref[i-1]
	}

	suff[len(nums)-1] = 1
	for i := len(nums)-2; i >= 0 ; i-- {
		suff[i] = nums[i+1] * suff[i+1] 
	}

	for i := 0; i < len(nums); i++ {
		res = append(res, pref[i] * suff[i])
	}

	return res
}
