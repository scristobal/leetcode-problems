func myAtoi(s string) int {
    
    r := 0

    i := 0
    for i<len(s) {
        if s[i] != byte(' ') { break }
        i++
    }
    
    if i>=len(s) { return  0 }
    
    pos := s[i] != byte('-') 
    
    if s[i] == byte('-') || s[i] == byte('+'){
        i++        
    }
    
    for i<len(s) {
        if s[i]<byte('0') || byte('9')< s[i] { break }
        
        d := s[i] - byte('0') 
        r = r*10 + int(d)
        
        if r > math.MaxInt32 { 
            r = math.MaxInt32
            if !pos { r = r+1  }
            break
        }
        
        i++
    }
    
    if !pos { return -r }
    return r
}