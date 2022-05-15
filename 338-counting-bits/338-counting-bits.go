func countBits(n int) []int {
    res := make([]int, n+1)
    
    for j := range res {
        res[j] = j % 2  + res[j >> 1]
    }
    
    return res
}