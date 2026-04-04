func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	for i := 0; i < len(s2) - len(s1) - 1; i ++ {
		if isPermutation(s1, s2[i:i+len(s1)]) {
			return true
		}
	}
	
	return false
}


func isPermutation(s1, s2 string) bool {
	mp := make(map[rune]int)
	for _, c := range s1 {
		mp[c]++
	}

	for _, c := range s2 {
		mp[c]--
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	return true
}