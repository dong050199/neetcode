func lengthOfLongestSubstring(s string) int {
	// now we need to solve it with O(n)
	left, right := 0, 1

	res := 0

	mp := make(map[byte]int)// character and indice.
	mp[s[left]] = 0
 
	for l < r {
		if i, exist mp[s[right]]; !exist {
			res = max(res, right - left + 1)
			mp[s[right]] = i // set indice
		} else {
			// left now become the next of the found of right value
			left = mp[s[right]] + 1



			// where is the new left pointer
		}
		// alway move right
		right++
	}




	return res
}
