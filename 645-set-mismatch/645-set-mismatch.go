func findErrorNums(nums []int) []int {
    
    l:= len(nums)
    
    occs := make([]int, l)
    
    for i := 0; i < l; i++ {
        
        occs[nums[i]-1] +=1
    }
    
    var miss int
    var dupl int
    
    for i := 0; i < l; i++ {
        switch n := occs[i]; {
            case n == 0 : miss = i + 1;
            case n > 1: dupl = i + 1;
        }
        
    }
      
    return []int{dupl,miss}
}