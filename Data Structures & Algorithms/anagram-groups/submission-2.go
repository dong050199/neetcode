func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	mp := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int
		for i := 0; i < len(str); i++ {
			key[str[i] - 'a']++ 
		}
		mp[key] = append(mp[key], str)
	}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}
