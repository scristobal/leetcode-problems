func pivotIndex(nums []int) int {

    right := 0
    
    for _, v := range nums{
        right += v 
    }

    
    left := 0
    
    for ind, val := range nums {
        
        right -= val
        
        if left == right { return ind }
        
        //if left > right { return -1 }
        
        left += val
    }
    
    return -1 
}