func minOperations(boxes string) []int {
    
    l := len(boxes)
    
    left := make([]int, l)
    
    for i, acc := 1, 0; i < l ; i++ {
        if boxes[i-1] == byte(49) { acc++ }
        left[i] = left[i-1] + acc
    }
    
    
    right := make([]int, l)

    for i, acc := l-2, 0; i>=0; i-- {
        if boxes[i+1] == byte(49) { acc++ }
        right[i] = right[i+1] + acc
    }
    
    res := make([]int, l)
    
    for i:=0;i<l;i++{
        res[i]= left[i] + right[i]
    }
    
    
    return res
    
}