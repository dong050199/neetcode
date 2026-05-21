func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	mp := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		mp[s1[i]]++
	}

	l, r := 0, len(s1)

	for r <= len(s2) {
		if isPermutation(s1, s2[l:r]) {
			return true
		}
		l++
		r++
	}

	return false
}

func isPermutation(s1, s2 string) bool {
	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		arr1[s1[i] - 'a']++
		arr2[s2[i] - 'a']++
	}

	return arr1 == arr2
}
