func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := merge(nums1, nums2)
	l := len(mergedArr)
	if l % 2 != 0 {
		return float64(mergedArr[l/2])
	}
	res := (float64(mergedArr[l/2]) + float64(mergedArr[l/2 - 1])) / 2
	return  res
}


// this cound be O(n)
func merge(nums1, nums2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i ++
		} else {
			res = append(res, nums2[j])
			j ++
		}
	}

	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}

	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}

	return res
}
