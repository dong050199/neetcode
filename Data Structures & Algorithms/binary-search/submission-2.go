func search(nums []int, target int) int {
    l, r := 0, len(nums) -1

    for l < r {
        m := l + (r - l)/2
        if nums[m] == target {
            return m
        }
        if nums[m] > target {
            r = m -1 
            continue
        }
        l = m 
    }
    return -1
}

// because array was sorted and solution must run on O(log n) time so
// binary search is the most efficent to run this.
// what is the binary search, it work by repeatedly cutting the search space in half
// if target is larger -> search for left
// if taeget is smaller -> search for right

