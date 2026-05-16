func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
    mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	l := 0
	curMin := 1001
	curRes := ""
	for r := 0; r < len(s); r++ {
		if _, exist := mp[s[r]]; exist {
			mp[s[r]]--
		}

		for isAllValueNonPositive(mp) {
			if _, exist := mp[s[l]]; exist {
				mp[s[l]]++
			}

			if r - l + 1 < curMin {
				curRes = s[l: r+1]
				curMin = r - l + 1
			}
			l++
		}
	}

	return curRes
}

func isAllValueNonPositive(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}

	return true
}
