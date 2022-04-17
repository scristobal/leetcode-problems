


func moveZeroes(nums []int)  {
    
    //fmt.Println("---")
    //fmt.Println("input: ", nums)

    l := len(nums) - 1
    
    if l==0 {return}
    
    //swap := reflect.Swapper(nums)
    
    
    i,j := 0,0

    for j < l {
                
        for i < l  && nums[i] != 0  {
            i++
        }
        
        j = i
        for j < l && nums[j] == 0 {
             j++ 
        }
       
        // swap(i,j)
        nums[i], nums[j] = nums[j], nums[i] 
        
        //fmt.Println(nums, i, j )
    }
    
    
}