func pivotIndex(nums []int) int {

    right := 0
    
    for _, v := range nums{
        right += v 
    }

    
    for i, left :=0, 0; i < len(nums); i++ {
        
        right -= nums[i]
        
        if left == right { return i }
        
        //if left > right { return -1 }
        
        left += nums[i]
    }
    
    return -1 
}