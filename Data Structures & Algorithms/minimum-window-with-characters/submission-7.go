func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	res := ""
	minLen := len(s) + 1

    mp := make(map[byte]int) 
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	l := 0

	for i := 0; i < len(s); i++ {
		if _, exist := mp[s[i]]; exist {
			mp[s[i]]--
		}

		for allValNonPositive(mp) {
			if _, exist := mp[s[l]]; !exist {
				l++
				continue
			}
			
			mp[s[l]]++
			if i - l + 1 <= minLen {
				minLen = i - l + 1
				res = s[l:i+1]
			}

			l++
		}
	}
	return res
}

func allValNonPositive(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}

	return true
}
