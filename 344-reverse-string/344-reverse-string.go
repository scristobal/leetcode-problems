func reverseString(s []byte)  {
    n := len(s);
    
    i := 0
    
    for i < n/2 {
        s[i], s[n-i-1] = s[n-i-1], s[i]
        
        i ++
    }
   
}