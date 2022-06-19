func rearrangeArray(nums []int) []int {
   
    res := make([]int, len(nums))
    
    p, n, l:= 0, 0, -1
    
    for p < len(nums) && nums[p] < 0 { p++ }
    for n<len(nums) && nums[n] > 0 { n++ }
    
    for j:=0; j < len(nums); j++ {
        
        if l < 0 { 
            res[j] = nums[p]
            l = nums[p]
            p++
            for p < len(nums) && nums[p] < 0 { p++ }
        } else {
            res[j] = nums[n]
            l = nums[n]
            n++
            for n<len(nums) && nums[n] > 0 { n++ }
        }
    }
    
    return res
    
}