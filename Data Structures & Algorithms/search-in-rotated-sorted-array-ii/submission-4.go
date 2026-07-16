func search(nums []int, target int) bool {
    l, r := 0, len(nums)-1

    for l <= r {
        mid := l + (r-l)/2

        if nums[mid] == target {
            return true
        }

        // cannot determine the sorted half
        if nums[l] == nums[mid] && nums[mid] == nums[r] {
            l++
            r--
        } else if nums[l] <= nums[mid] { // left half sorted
            if nums[l] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else { // right half sorted
            if nums[mid] < target && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }

    return false
}