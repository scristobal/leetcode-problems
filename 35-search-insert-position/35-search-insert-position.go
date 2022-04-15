func searchInsert(nums []int, target int) int {
    lo := 0
    hi := len(nums) 
    
    // fmt.Println("---")
    
    for lo<hi {
        mid := lo + (hi-lo)/2;
        
        //if nums[mid] == target { return mid }
        
        if  nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid 
        }
       
        // fmt.Printf("lo = %v, hi = %v\n", lo, hi)
        
    }
    
    // fmt.Printf("lo = %v, hi = %v\n", lo, hi)
    
    return lo
}