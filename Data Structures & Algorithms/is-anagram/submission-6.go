func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	str1 := [26]int{}
	str2 := [26]int{}

	for i := 0; i < len(s); i++ {
		str1[s[i] - 'a']++
		str2[t[i] - 'a']++
	}

	return str1 == str2
}
