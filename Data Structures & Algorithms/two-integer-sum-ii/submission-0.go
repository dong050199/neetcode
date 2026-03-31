func twoSum(numbers []int, target int) []int {
    i := 0 // first element
    j := len(numbers) - 1 // last element
    for {
        sum := numbers[i] + numbers[j] 
        if sum > target {
            j--
        }

        if sum < target {
            i++
        }

        if sum == target {
            return []int{i+1,j+1}
        }
    }
}

// using two pointer
// this also can be solve by using hash map