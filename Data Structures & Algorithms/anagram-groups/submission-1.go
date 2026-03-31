func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	var res [][]string

	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c - 'a']++
		}

		mp[key] = append(mp[key], str)
	}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}
