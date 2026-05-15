func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	minLen := len(s)
	res := ""

	l := 0
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			mp[s[r]]--
		}

		for isAllValueNonPositive(mp) {
			if _, exist := mp[s[l]]; !exist {
				l++
				continue
			}

			mp[s[l]]++

			if (r - l + 1) <= minLen {
				minLen = r - l + 1
				res = s[l:r+1]
			}

			l++
		}
	}
	return res
}

func isAllValueNonPositive(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}

	return true
}
