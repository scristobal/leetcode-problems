func isUgly(n int) bool {
    
    if n <= 0 { return false }
    if n == 1  { return true }
    
    r := n;
    
    for r % 2 == 0 {
        r = r/2;
    }
    
    for r % 3 == 0 {
        r = r/3;
    }
    
    for r % 5 == 0 {
        r = r/5;
    }
    
    if r == 1 { return true }
    return false
}