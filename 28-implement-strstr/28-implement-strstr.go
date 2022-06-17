func strStr(haystack string, needle string) int {
    
    l := len(needle) 
    
    if l==0 {return 0}
    
    for h := range haystack {
        if h + l > len(haystack) {break }
        if needle == haystack[h:h+l] { return h }
    }
    
    return -1
}