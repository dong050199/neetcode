func minWindow(s string, t string) string {
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	res := ""
	l := 0
	for r := 0; r < len(s); r++ {  // ✅ Fixed: r++, not i++
		// Expand phase
		if _, exist := mp[s[r]]; exist {
			mp[s[r]]--  // Decrement, DON'T delete
		}

		// Shrink phase - while all needed chars are found
		for allNonPositive(mp) {
			// Record window
			if res == "" || r-l+1 < len(res) {
				res = s[l : r+1]
			}
			
			// Try to shrink from left
			if _, exist := mp[s[l]]; exist {
				mp[s[l]]++  // Increment back
			}
			l++
		}
	}
	return res
}

func allNonPositive(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}
	return true
}