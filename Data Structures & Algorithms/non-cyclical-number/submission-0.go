func isHappy(n int) bool {
	mp := make(map[int]bool)

	for {
		out := sumOfSquare(n)
		if out == 1 {
			return true
		}

		if mp[out] {
			return false
		}

		mp[out] = true
		n = out
	}

	return false
}

func sumOfSquare(n int) int {
	output := 0

	for n > 0 {
		digit := n % 10
		digit = digit * digit
		output += digit
		n = n / 10
	}
	return output
}
