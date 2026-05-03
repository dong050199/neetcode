func twoSum(numbers []int, target int) []int {
	res := []int{}
	l, r := 0, len(numbers) - 1
	for l < r {
		if numbers[l] + numbers[r] == target {
			return []int{l+1, r+1}
		}

		if numbers[l] + numbers[r] >= target {
			r--
			continue
		}

		if numbers[l] + numbers[r] < target {
			l++
			continue
		}
	}

	return res
}

// [1,2,3,4] - target = 3
// l -----r
// arr[l] + arr[r] == 3 => res
// arr[l] + arr[r] > 3 -> r--
// arr[l] + arr[r] < 3 -> l++