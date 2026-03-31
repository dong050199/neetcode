func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int32][]string)
	for _, str := range strs {
		m[getCharactersFrequency(str)] = append(m[getCharactersFrequency(str)], str)
	}

	output := make([][]string, 0)

	for _, v := range m {
		output = append(output, v)
	}

	return output
}

func getCharactersFrequency(str1 string) [26]int32 {
	var counter [26]int32
	for _, c := range str1 {
		counter[c-'a']++
	}

	return counter
}