func rearrangeArray(nums []int) []int {
   
    res := make([]int, len(nums))
    
    p, n:= 0, 1
    
    
    for j:=0; j < len(nums); j++ {
        
        if nums[j] > 0 {
            res[p] = nums[j]
            p +=2
            continue
        }
        
        res[n] = nums[j]
        n +=2
        
    }
    
    return res
    
}