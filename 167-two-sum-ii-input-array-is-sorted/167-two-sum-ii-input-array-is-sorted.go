func twoSum(nums []int, target int) []int {
    var visited = make(map[int]int)
    
    for i, v := range nums {
       
        
        j, ok := visited[target - v]
        
        if ok { 
            return []int{j+1,i+1} 
        }
        
    
        
        
        visited[v] = i
        
        for _, w := range visited {
            if w + v > target { delete(visited, w)}
        }
       
        
    }
    
    return []int{0,1}
}