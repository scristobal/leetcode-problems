func numTrees(n int) int {
    nBSTs := make([]int, n+1)
    
    nBSTs[0] = 1
    nBSTs[1] = 1
    
    for i := 2; i<len(nBSTs); i++{
       
        for j := i-1; j>=0; j-- {
            nBSTs[i] += nBSTs[(i-1)-j] * nBSTs[j]
        }
        
    }
    
    return nBSTs[n]
}