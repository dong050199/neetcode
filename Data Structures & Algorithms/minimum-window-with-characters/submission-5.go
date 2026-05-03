func minWindow(s string, t string) string {
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	l := 0
	res := ""
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			mp[s[r]]--
		}

		for isValidMap(mp) {
			tmp := s[l:r+1]
			if res == "" || len(tmp) < len(res) {
				res = tmp
			}

			if _, exist := mp[s[l]]; exist {
				mp[s[l]]++
			}

			l++
		}
	}
	return res
}

func isValidMap(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}
	return true
}
