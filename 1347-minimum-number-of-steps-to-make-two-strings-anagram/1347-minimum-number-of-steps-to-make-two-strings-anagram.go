func minSteps(s string, t string) int {
    count := make(map[byte]int)
    
    for i := 0; i < len(t); i++{
        count[s[i]] +=1
        count[t[i]] -=1
    }
    
    c := 0
    
    for _, v := range count {
        
        if (v >0) { c +=v } 
    }
    
    return c
}