func buildArray(nums []int) []int {
    
    res := make([]int, len(nums))
    
    for i, n := range nums {
        res[i] = nums[n]
    }
    
    return res
}

