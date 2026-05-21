func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	maxf := 0
	res := 0

	l := 0
	for r := 0; r < len(s); r++ {
		freqMap[s[r]]++
		maxf = max(maxf, freqMap[s[r]])

		for (r - l + 1) - k > maxf {
			freqMap[s[l]]--
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}