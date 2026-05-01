func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{numbers[left], numbers[right]}
		}

		if numbers[left] + numbers[right] >= target {
			right--
			continue
		}

		if numbers[left] + numbers[right] < target {
			left++
			continue
		}
	}

	return []int{}
}
