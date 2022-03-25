func firstUniqChar(s string) (res int) {
    
    res =  -1
    
    seenOnceAt := make(map[rune]int, len(s))
    seenMany := make(map[rune]bool, len(s))
    
    for i, c := range s {    
        
        _, seen := seenOnceAt[c] 
        
        if seen { 
            seenMany[c] = true;
            delete(seenOnceAt, c)
            continue
        }
        
        if seenMany[c] { 
            continue
        }
        
        seenOnceAt[c] = i
        
      
    }
    
    
    m := len(s)
    
    if len(seenOnceAt) == 0 { return -1}
    
    for _, i := range seenOnceAt {
        if i < m { m = i }
    }
    
    return m
}