func removeElement(nums []int, val int) int {

    j:=0;
    
    for _, n := range nums {
        if n != val {
            nums[j] = n
            j++
        }  
    }
    
    return j
}