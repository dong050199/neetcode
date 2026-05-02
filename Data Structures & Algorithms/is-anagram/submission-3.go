func isAnagram(s string, t string) bool {
	mp := make(map[rune]int)

	for _, c := range s {
		mp[c]++
	}

	for _, c := range t {
		mp[c]--
		if mp[c] == 0 {
			delete(mp, c)
		}
	}

	return len(mp) == 0
}
