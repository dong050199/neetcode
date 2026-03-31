func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp := make(map[rune][]string)
	mp['2'] = []string{"a", "b", "c"}
	mp['3'] = []string{"d", "e", "f"}
	mp['4'] = []string{"g", "h", "i"}
	mp['5'] = []string{"j", "k", "l"}
	mp['6'] = []string{"m", "n", "o"}
	mp['7'] = []string{"p", "q", "r", "s"}
	mp['8'] = []string{"t", "u", "v"}
	mp['9'] = []string{"w", "x", "y", "z"}
	var input [][]string
	for _, d := range digits {
		input = append(input, mp[d])
	}

	var res []string 
	dfs(input, []string{}, 0, &res)
	return res
}

func dfs(input [][]string, cur []string, i int, res *[]string) {
	if len(cur) == len(input) {
		str := ""
		for _, c := range cur {
			str = str + c	
		}	
		*res = append(*res, str)
		return 
	}
	
	for _, c := range input[i] {
		cur = append(cur, c)
		dfs(input, cur , i+1, res)
		cur = cur[:len(cur)-1]
	}

	return
}
