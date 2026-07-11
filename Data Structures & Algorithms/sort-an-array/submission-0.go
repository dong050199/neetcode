func sortArray(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }

    mid := len(nums) / 2
    left := sortArray(nums[:mid])
    right := sortArray(nums[mid:])
    return mergeSort(left, right)
}

func mergeSort(left, right []int) []int {
    res := []int{}
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            res = append(res, left[i])
            i++
        } else {
            res = append(res, right[j])
            j++
        }
    }
    res = append(res, left[i:]...)
    res = append(res, right[j:]...)
    return res
}