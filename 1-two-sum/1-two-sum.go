func twoSum(nums []int, target int) []int {
    var visited = make(map[int]int)
    
    for i, v := range nums {
        
        j, ok := visited[target - v]
        
        if ok { 
            return []int{i, j} 
        }
        
        visited[v] = i
        
    }
    
    return []int{0,1}
}