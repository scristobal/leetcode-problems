func searchInsert(nums []int, target int) int {
    lo := 0
    hi := len(nums) 
    
    fmt.Println("---")
    
    for (hi-lo)>1 {
        mid := lo+ (hi-lo)/2;
        
        if nums[mid] == target { return mid }
        
        if  nums[mid] < target {
            lo = mid 
        } else {
            hi = mid 
        }
       
        fmt.Printf("lo = %v, hi = %v\n", lo, hi)
        
    }
    
    if target <= nums[lo] { return lo}
    
    return hi 
}