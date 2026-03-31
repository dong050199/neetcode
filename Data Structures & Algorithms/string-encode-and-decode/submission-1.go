type Solution struct{
	order []int
}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "*"
		res += str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
    var res []string
    for len(encoded) > 0 {
        i := strings.Index(encoded, "*")
        num, _ := strconv.Atoi(encoded[:i])
        encoded = encoded[i+1:]      
        res = append(res, encoded[:num])
        encoded = encoded[num:]
    }
    return res
}
