func maximumProduct(nums []int) int {
    var maxs = [3]int{math.MinInt, math.MinInt, math.MinInt}
    var mins = [2]int{math.MaxInt, math.MaxInt}
    
    for _, v := range nums {
        switch {
            case v > maxs[2]: 
                maxs[0] = maxs[1]
                maxs[1] = maxs[2]
                maxs[2] = v
            case v > maxs[1]:
                maxs[0] = maxs[1]
                maxs[1] = v
            case v > maxs[0]:
                maxs[0] = v
                
        }
        
        switch {
            case v < mins[1]:
                mins[0] = mins[1]
                mins[1] = v
            case v < mins[0]:
                mins[0] = v
        }
    }
    
  
    
    min := mins[0]*mins[1]*maxs[2]
    max := maxs[0]*maxs[1]*maxs[2]
    
    if min > max  {
        return min
    }
    
    return max
}