func search(ns []int, t int) int {
    l := (len(ns))
    
    if l == 0 { return -1}
    
    s := math.Ceil(float64(l/2))
    i := int(s)
    
    for {
        
        if ns[i] == t { return i}
           
        s = math.Ceil(s/2)
        
        if s == 1 || s == 0 { 
            if i-1>= 0 && ns[i-1] == t { return i -1}
            if i +1 < l && ns[i + 1] == t { return i +1 }
            return -1 
        }
        
        
        if ns[i]<t { 
            i += int(s)
            continue
        }
        
        i -= int(s)
    }
    
    return -1
    
}