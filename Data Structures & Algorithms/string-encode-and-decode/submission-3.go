type Solution struct{
    length []int
}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        res+= str
        s.length = append(s.length, len(str))
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    for _, length := range s.length {
        res = append(res, encoded[:length])
        encoded = encoded[length:]
    }
    return res
}
