func permute(nums []int) [][]int {
	var res [][]int
	backtracking(nums, []int{}, make([]bool, len(nums)), &res)
	return res
}

func backtracking(nums []int, cur []int, pick []bool, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := append([]int{}, cur...)
		*res = append(*res, tmp)
		return
	}

	for i:= 0; i< len(nums); i++ {
		if !pick[i] {
			pick[i] = true // mark an number was picked to not go thought it again.
			cur = append(cur, nums[i]) // append until we reach the len of nums.
			backtracking(nums, cur, pick, res)
			// back tracking logic here.
			cur = cur[:len(cur)-1]
			pick[i] = false
		}
	}
}
