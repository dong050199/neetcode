func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		arr1[s[i] - 'a']++
		arr2[t[i] - 'a']++
	}

	return arr1 == arr2
}
