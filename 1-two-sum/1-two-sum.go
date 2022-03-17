func twoSum(nums []int, target int) []int {
    var visited = make(map[int]int)
    
    for i := range nums {
        
        j, ok := visited[target - nums[i]]
        
        if ok { 
            return []int{i, j} 
        }
        
        visited[nums[i]] = i
        
    }
    
    return []int{0,1}
}