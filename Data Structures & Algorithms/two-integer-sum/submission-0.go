func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, num := range nums {
        m[num] = i
    }

    for i := 0; i < len(nums); i ++ {
        diff := target - nums[i]
        if v, exist := m[diff]; exist {
            return []int{i,v}
        }
    }

    return []int{}
}

// 3,4,5,6
// 0:4 
// 1:3 
// 2:2 
// 3:1
// 
