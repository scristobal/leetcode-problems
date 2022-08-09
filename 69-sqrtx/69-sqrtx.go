func mySqrt(x int) int {

   
    
    r := 0
    
    for r*r < x {
        r++
    }
    
    if (r*r == x) { return r}
    
    return r-1
}