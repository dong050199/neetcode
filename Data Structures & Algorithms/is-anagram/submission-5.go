func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	arrS, arrT := [26]int{}, [26]int{}
	for i :=0 ; i < len(s); i++ {
		arrS[s[i] - 'a']++
		arrT[t[i] - 'a']++
	}

	return arrS == arrT
}
