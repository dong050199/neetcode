func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		arr := [26]int{}
		for i := 0; i < len(str); i++ {
			arr[str[i] - 'a']++
		}

		mp[arr] = append(mp[arr], str)
	}

	res := [][]string{}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}
