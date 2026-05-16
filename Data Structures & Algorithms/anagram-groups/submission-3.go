func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		cur := [26]int{}
		for i := 0; i < len(str); i++ {
			cur[str[i] - 'a']++
		}

		mp[cur] = append(mp[cur], str)
	}

	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}

	return res
}
