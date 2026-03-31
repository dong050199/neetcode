func isAnagram(s string, t string) bool {
	mp := make(map[rune]int)
	for _, c := range s {
		mp[c]++
	}

	for _, c := range t {
		mp[c]-- 
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	return true
}
