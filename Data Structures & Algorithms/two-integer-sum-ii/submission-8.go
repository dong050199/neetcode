func twoSum(numbers []int, target int) []int {
	l,r := 0l len(nums)-1

	for l < r {
		if numbers[l] + numbers[r] == target {
			return []int{l, r}
		}

		if numbers[l] + numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	return []int{}
}
